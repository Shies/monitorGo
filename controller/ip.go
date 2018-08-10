package controller

import (
	"html/template"

	"monitorGo/dao"
	"monitorGo/model"
)

func ipList(c Context) {
	var (
		req = c.Request()
		sql   string
		param int64
	)
	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		param = 999
		sql = dao.IPS_BY_ALL
	} else {
		sql = dao.IPS_BY_TID
		param = parseInt(query["tid"][0])
	}
	tasks, _ := srv.TaskList(dao.TASK_BY_ALL, "1")
	ips, _ := srv.TaskIP(sql, param)
	res := map[string]interface{}{
		"Task": tasks,
		"Ips": ips,
	}
	c.SetData(res)
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
