package net

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"log"
	"net"
	"time"
	"crypto/tls"
	"monitorGo/conf"
)

type Client struct {
	conf	  *conf.HttpClient
	client    *http.Client
	dialer    *net.Dialer
	transport http.RoundTripper
}

func NewClient(c *conf.HttpClient) *Client {
	client := new(Client)
	client.conf = c
	client.dialer = &net.Dialer{
		Timeout:   c.Timeout * time.Second,
		KeepAlive: c.KeepAlive * time.Second,
	}
	client.transport = &http.Transport{
		DialContext:     client.dialer.DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.client = &http.Client{
		Transport: client.transport,
	}

	return client
}

func (c *Client) HttpDo(method string, requestUrl string, params string, header map[string]string) (string, error) {
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
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if header != nil {
		req.Header.Set("Host", header["host"])
	}

	resp, err := c.client.Do(req)
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
func (c *Client) ParseUrl(requestUrl string, ip string) map[string]string {
	var (
		path string
		part = make(map[string]string)
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
	part["url"] = urlData.Scheme + "://" + ip + path
	part["header"] = "Host:" + urlData.Host

	return part
}