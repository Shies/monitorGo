package controller

import (
	"html/template"

	"monitorGo/dao"
	"monitorGo/model"
)

func noticeList(c Context) {
	var (
		param int64
		sql   string
		req = c.Request()
	)
	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		param = 999
		sql = dao.NOTICE_BY_ALL
	} else {
		sql = dao.NOITCE_BY_TID
		param = parseInt(query["tid"][0])
	}
	resp := make(map[string]interface{})
	resp["Notice"] = srv.SendList(sql, param)
	resp["Task"] = srv.TaskList(dao.TASK_BY_ALL, "1")
	c.SetData(resp)
	c.SetPath("views/notice.html")
	views(c)
}

func saveNotice(c Context) {
	var (
		req = c.Request()
		res = c.Response()
	)
	req.ParseForm()
	for k, v := range req.Form {
		req.Form[k][0] = template.HTMLEscapeString(v[0])
	}
	notice := &model.Notice{
		SendType: atoi(req.PostFormValue("newsendtype")),
		Content:  req.PostFormValue("newcontent"),
		Tid:      parseInt(req.PostFormValue("newtid")),
	}
	srv.SaveNotice(notice)
	// 跳转
	res.Header().Add("Location", "/noticeTpl")
	res.WriteHeader(302)
}
