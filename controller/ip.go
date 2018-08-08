package controller

import (
	"html/template"

	dao2 "monitorGo/dao"
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
		sql = dao2.IPS_BY_ALL
	} else {
		sql = dao2.IPS_BY_TID
		param = parseInt(query["tid"][0])
	}
	resp := make(map[string]interface{})
	resp["Task"] = dao.TaskList(dao2.TASK_BY_ALL, "1")
	resp["Ips"] = dao.TaskIP(sql, param)
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
	dao.SaveIP(ip)
	// 跳转
	res.Header().Add("Location", "/ipTpl")
	res.WriteHeader(302)
}
