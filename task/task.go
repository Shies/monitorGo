package task

import (
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"log"

	"monitorGo/dao"
	"monitorGo/conf"
)


var (
	_ = conf.ParseConfig()
	d = dao.New(&conf.Conf)
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

func SendMails(emails []string, msg string) error {
	var (
		err error
		mail = new(Mail)
	)
	confs, _ := d.ConfList()
	globalEmails := strings.Split(confs["GlobalEmails"], ",")
	for _, email := range globalEmails {
		emails = append(emails, email)
	}

	log.Printf("%v\n", emails)
	for _, email := range emails {
		err = mail.Send(email, "告警邮件", msg, "normal")
		if err != nil {
			log.Printf("%v\n", err)
		}
	}
	return err
}

func HttpDo(method string, requestUrl string, params string, header map[string]string) (string, error) {
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
func ParseUrl(requestUrl string, ip string) map[string]string {
	var (
		path string
		ret  = make(map[string]string)
	)

	urlData, err := url.Parse(requestUrl)
	if err != nil {
		log.Println(err)
		return nil
	}
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