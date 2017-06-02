package controller

import (
	"html/template"
	"net/http"
	"strings"
)

func ConfTpl(w http.ResponseWriter, req *http.Request) {
	conf := dao.GetConf()
	views("./views/conf.html", conf, w)
}

func SaveConf(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for key, val := range req.Form {
		_val := template.HTMLEscapeString(strings.Join(val, ""))
		dao.SaveConf(key, _val)
	}

	w.Header().Add("Location", "/confTpl")
	w.WriteHeader(302)
}
