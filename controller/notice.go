package controller

import (
	"html/template"
	"monitorGo/model"
	"net/http"
)

func noteList(w http.ResponseWriter, req *http.Request) {
	var (
		param int64
		sql   string
	)

	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		param = 999
		sql = model.NOTICE_BY_ALL
	} else {
		sql = model.NOITCE_BY_TID
		param = parseInt(query["tid"][0])
	}

	resp := make(map[string]interface{})
	resp["Notice"] = dao.SendList(sql, param)
	resp["Task"] = dao.TaskList(model.TASK_BY_ALL, "1")

	views("views/notice.html", resp, w)
}

func saveNote(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for k, v := range req.Form {
		req.Form[k][0] = template.HTMLEscapeString(v[0])
	}

	notice := &model.Notice{
		SendType: atoi(req.PostFormValue("newsendtype")),
		Content:  req.PostFormValue("newcontent"),
		Tid:      parseInt(req.PostFormValue("newtid")),
	}

	dao.SaveNotice(notice)
	w.Header().Add("Location", "/noticeTpl")
	w.WriteHeader(302)
}
