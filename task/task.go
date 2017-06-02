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

const (
	_SMTP_USER   = "SmtpUser"
	_SMTP_PASS   = "SmtpPass"
	_SMTP_SERVER = "SmtpServer"
)

// Smtp struct info.
type Mail struct {
	user   string
	pass   string
	server string
}

func Configure(user string, pass string, server string) *Mail {
	mail := &Mail{
		user:   user,
		pass:   pass,
		server: server,
	}
	return mail
}

func (m *Mail) Send(to []string, subject string, body string) {
	mail := Configure(_SMTP_USER, _SMTP_PASS, _SMTP_SERVER)
	var host = strings.Split(mail.server, ":")
	// go func() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", mail.user, mail.pass, host[0])
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := []byte(
		"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")
	err := smtp.SendMail(mail.server, auth, mail.user, to, msg)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
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

	var path string
	_return := make(map[string]string)
	if urlData.Path != "" {
		path = urlData.Path
	}
	if urlData.RawQuery != "" {
		path = path + "?" + urlData.RawQuery
	}

	_return["header"] = "Host:" + urlData.Host
	_return["url"] = urlData.Scheme + "://" + ip + path

	return _return
}

func sendMails(emails []string, msg string) {
	mail := &Mail{}
	conf := model.New().GetConf()
	globalEmails := strings.Split(conf["GlobalEmails"], ",")
	mail.Send(emails, "告警邮件", msg)
	mail.Send(globalEmails, "告警邮件", msg)
}

func Request(t *model.TaskItem, ips []*model.TaskIP) {
	// time.Sleep(time.Duration(t.Frequency) * time.Minute)
	if ips == nil {
		resp := httpDo(t.Method, t.Url, t.Params, nil)
		if resp == "" {
			return
		}
	}

	header := make(map[string]string)
	for _, v := range ips {
		urlData := parseUrl(t.Url, v.IP)
		part := strings.Split(urlData["header"], ":")
		header["host"] = part[1]
		resp := httpDo(t.Method, urlData["url"], t.Params, header)
		if resp == "" {
			return
		}
	}
}
