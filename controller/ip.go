package controller

import (
	"html/template"

	"monitorGo/dao"
	"monitorGo/model"
)

func ipList(c Context) {
	var (
		req = c.Request()
		param int64
		sql   string
	)
	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		param = 999
		sql = dao.IPS_BY_ALL
	} else {
		sql = dao.IPS_BY_TID
		param = parseInt(query["tid"][0])
	}
	resp := make(map[string]interface{})
	resp["Task"] = srv.TaskList(dao.TASK_BY_ALL, "1")
	resp["Ips"] = srv.TaskIP(sql, param)
	c.SetData(resp)
	c.SetPath("views/ip.html")
	views(c)
}

func saveIP(c Context) {
	var (
		req = c.Request()
		res = c.Response()
	)
	req.ParseForm()
	ip := &model.TaskIP{
		Tid: parseInt(req.PostFormValue("tid")),
		IP:  template.HTMLEscapeString(req.PostFormValue("ips")),
	}
	srv.SaveIP(ip)
	// 跳转
	res.Header().Add("Location", "/ipTpl")
	res.WriteHeader(302)
}
