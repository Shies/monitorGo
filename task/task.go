package task

import (
	"fmt"
	"io/ioutil"
	"monitorGo/model"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
)

var (
	dao = model.New()
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

func httpDo(method string, requestUrl string, params string, header map[string]string) string {
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
		return ""
	}
	if header != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Host", header["host"])
	}

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	return string(body)
}

// parse url and set it's ip
func parseUrl(requestUrl string, ip string) map[string]string {
	urlData, err := url.Parse(requestUrl)
	if err != nil {
		fmt.Println(err)
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
	mail := &Mail{}

	conf := dao.ConfList()
	globalEmails := strings.Split(conf["GlobalEmails"], ",")
	for _, email := range globalEmails {
		emails = append(emails, email)
	}

	fmt.Printf("%v\n", emails)
	for _, email := range emails {
		err := mail.Send(email, "告警邮件", msg, "normal")
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}

	return true
}

func Request(t *model.TaskItem, ips []*model.TaskIP) bool {
	// time.Sleep(time.Duration(t.Frequency) * time.Minute)
	if ips == nil {
		resp := httpDo(t.Method, t.Url, t.Params, nil)
		if resp == "" {
			fmt.Printf("%s\n", t.Url)
			return false
		}
	}

	header := make(map[string]string)
	for _, v := range ips {
		urlData := parseUrl(t.Url, v.IP)
		part := strings.Split(urlData["header"], ":")
		header["host"] = part[1]
		go func() {
			resp := httpDo(t.Method, urlData["url"], t.Params, header)
			if resp == "" {
				fmt.Printf("%s\n", urlData["url"])
			}
		}()
	}

	return true
}
