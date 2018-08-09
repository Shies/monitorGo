package task

import (
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"sync"
	"log"
	"fmt"

	"monitorGo/conf"
	"monitorGo/model"
	"monitorGo/dao"
	"time"
)


var (
	d = dao.New()
	_ = conf.Logger(conf.Conf.Log.Dir)
)

const (
	_SMTP_USER   = "your email username"
	_SMTP_PASS   = "your email password"
	_SMTP_SERVER = "your email server"
)

// Smtp struct info.
type Mail struct {
	user   string
	pass   string
	server string
}

func (m *Mail) Send(to, subject, body, mailtype string) error {
	mail := &Mail{
		user:   _SMTP_USER,
		pass:   _SMTP_PASS,
		server: _SMTP_SERVER,
	}

	hp := strings.Split(mail.server, ":")
	auth := smtp.PlainAuth("", mail.user, mail.pass, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + mail.user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(mail.server, auth, mail.user, send_to, msg)
	return err
}

func httpDo(method string, requestUrl string, params string, header map[string]string) (string, error) {
	client := &http.Client{}
	var (
		req *http.Request
		err error
	)
	if "GET" == method || method == "" {
		req, err = http.NewRequest("GET", requestUrl, nil)
	} else {
		req, err = http.NewRequest("POST", requestUrl, strings.NewReader(params))
	}
	if err != nil {
		return "", err
	}
	if header != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Host", header["host"])
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// parse url and set it's ip
func parseUrl(requestUrl string, ip string) map[string]string {
	urlData, err := url.Parse(requestUrl)
	if err != nil {
		log.Println(err)
		return nil
	}

	var (
		path string
		ret  = make(map[string]string)
	)
	if urlData.Path != "" {
		path = urlData.Path
	}
	if urlData.RawQuery != "" {
		path = path + "?" + urlData.RawQuery
	}

	ret["url"] = urlData.Scheme + "://" + ip + path
	ret["header"] = "Host:" + urlData.Host

	return ret
}

func SendMails(emails []string, msg string) bool {
	mail := new(Mail)

	config := d.ConfList()
	globalEmails := strings.Split(config["GlobalEmails"], ",")
	for _, email := range globalEmails {
		emails = append(emails, email)
	}

	log.Printf("%v\n", emails)
	for _, email := range emails {
		err := mail.Send(email, "告警邮件", msg, "normal")
		if err != nil {
			log.Printf("%v\n", err)
		}
	}

	return true
}

type Event struct {
	test chan string
	send chan []string
	done chan bool
	sync *sync.WaitGroup
}

func (e *Event) Tester() {
	e.test <- "hello world"
}

func (e *Event) Producer(ips []*model.TaskIP) {
	defer e.sync.Done()
	if ips != nil {
		var ipstr = []string{}
		for _, v := range ips {
			ipstr = append(ipstr, v.IP)
		}
		select {
		case e.send <- ipstr:
		default:
			log.Println("the chan is full(" + strings.Join(ipstr, ",") + ")")
		}
	}
}

func (e *Event) Consumer(t *model.TaskItem) {
	for {
		select {
		case ipstr, ok := <-e.send:
			if !ok {
				e.Close()
				return
			}
			var header = make(map[string]string)
			for _, ip := range ipstr {
				log.Println("start:" + t.Url)
				urlData := parseUrl(t.Url, ip)
				part := strings.Split(urlData["header"], ":")
				header["host"] = part[1]
				if _, err := httpDo(t.Method, urlData["url"], t.Params, header); err != nil {
					log.Printf("%v\n", urlData)
				}
			}
		case te := <-e.test:
			fmt.Println(te)
		case <-e.done:
			e.Close()
			log.Println("done")
			return
		}
	}
}

func (e *Event) Done() {
	time.Sleep(time.Duration(3) * time.Second)
	e.done <- true
}

func (e *Event) Close() {
	close(e.send)
	close(e.test)
	close(e.done)
}

func Request(t *model.TaskItem, ips []*model.TaskIP) {
	e := &Event{
		test: make(chan string),
		send: make(chan []string, len(ips)),
		done: make(chan bool, 1),
		sync: new(sync.WaitGroup),
	}

	if ips != nil {
		go e.Consumer(t)

		e.sync.Add(1)
		go e.Producer(ips)
		go e.Tester()
		go e.Done()
	} else {
		go func() {
			if _, err := httpDo(t.Method, t.Url, t.Params, nil); err != nil {
				log.Printf("%v\n", err)
				return
			}
		}()
	}
	return
}