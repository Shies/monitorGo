package controller

import (
	"strings"
	"html/template"
)

func confList(c Context) {
	conf, _ := srv.ConfList()
	c.SetData(conf)
	c.SetPath("views/conf.html")
	views(c)
}

func saveConf(c Context) {
	var (
		req = c.Request()
		res = c.Response()
	)
	req.ParseForm()
	for key, val := range req.Form {
		_val := template.HTMLEscapeString(strings.Join(val, ""))
		srv.SaveConf(key, _val)
	}
	res.Header().Add("Location", "/confTpl")
	res.WriteHeader(302)
}
