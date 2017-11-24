package controller

import (
	"html/template"
	"net/http"
	"strings"
)

func confList(w http.ResponseWriter, req *http.Request) {
	conf := dao.ConfList()
	views("views/conf.html", conf, w)
}

func saveConf(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for key, val := range req.Form {
		_val := template.HTMLEscapeString(strings.Join(val, ""))
		dao.SaveConf(key, _val)
	}

	w.Header().Add("Location", "/confList")
	w.WriteHeader(302)
}
