package controller

import (
	"strings"
	"net/http"
	"html/template"

	"monitorGo/model"
	"monitorGo/dao"
)

func taskList(c Context) {
	var (
		param string
		sql   string
		req = c.Request()
	)
	query := req.URL.Query()
	if len(query["name"]) == 0 {
		param = "1"
		sql = dao.TASK_BY_ALL
	} else if strings.Contains(query["name"][0], "http") {
		param = "'%" + query["name"][0] + "%'"
		sql = dao.TASK_BY_URL
	} else {
		param = "'%" + query["name"][0] + "%'"
		sql = dao.TASK_BY_NAME
	}
	task, _ := srv.TaskList(sql, param)
	c.SetData(task)
	c.SetPath("views/task.html")
	views(c)
}

func saveTask(c Context) {
	var (
		req = c.Request()
		res = c.Response()
	)
	req.ParseForm()
	for k, v := range req.Form {
		req.Form[k][0] = template.HTMLEscapeString(v[0])
	}
	taskItem := getTaskParam(req)
	if taskItem.Name == "" {
		panic("invalid params")
	}
	srv.SaveTask(taskItem)
	res.Header().Add("Location", "/taskTpl")
	res.WriteHeader(302)
}

func getTaskParam(req *http.Request) *model.TaskItem {
	taskItem := &model.TaskItem{}
	taskItem.Name = req.PostFormValue("name")
	taskItem.Protocol = req.PostFormValue("protocol")
	taskItem.Url = req.PostFormValue("url")
	taskItem.Username = req.PostFormValue("username")
	taskItem.Password = req.PostFormValue("password")
	taskItem.Method = req.PostFormValue("method")
	taskItem.Params = req.PostFormValue("params")
	taskItem.Frequency = atoi(req.PostFormValue("frequency"))
	taskItem.Retry = atoi(req.PostFormValue("retry"))
	taskItem.Goodcode = atoi(req.PostFormValue("goodcode"))
	taskItem.Status = atoi(req.PostFormValue("status"))
	taskItem.Uid = 1
	taskItem.Gid = 1
	taskItem.Sizerange = req.PostFormValue("sizerange")
	taskItem.Createtime = req.PostFormValue("createtime")
	taskItem.Respbody = req.PostFormValue("respbody")

	return taskItem
}
