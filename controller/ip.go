package controller

import (
	"html/template"
	"monitorGo/model"
	"net/http"
)

func ipList(w http.ResponseWriter, req *http.Request) {
	var (
		param int64
		sql   string
	)

	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		param = 999
		sql = model.IPS_BY_ALL
	} else {
		sql = model.IPS_BY_TID
		param = parseInt(query["tid"][0])
	}

	resp := make(map[string]interface{})
	resp["Task"] = dao.TaskList(model.TASK_BY_ALL, "1")
	resp["Ips"] = dao.TaskIP(sql, param)

	views("views/ip.html", resp, w)
}

func saveIP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	ip := &model.TaskIP{
		Tid: parseInt(req.PostFormValue("tid")),
		IP:  template.HTMLEscapeString(req.PostFormValue("ips")),
	}

	dao.SaveIP(ip)
	w.Header().Add("Location", "/ipList")
	w.WriteHeader(302)
}
