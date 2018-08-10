package controller

import (
	"html/template"

	"monitorGo/dao"
	"monitorGo/model"
)

func noticeList(c Context) {
	var (
		req = c.Request()
		sql   string
		param int64
	)
	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		param = 999
		sql = dao.NOTICE_BY_ALL
	} else {
		sql = dao.NOITCE_BY_TID
		param = parseInt(query["tid"][0])
	}
	send, _ := srv.SendList(sql, param)
	task, _ := srv.TaskList(dao.TASK_BY_ALL, "1")
	res := map[string]interface{}{
		"Notice": send,
		"Task": task,
	}
	c.SetData(res)
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
