package controller

import (
	"fmt"
	"html/template"
	"monitorGo/model"
	"net/http"
	"strings"
)

func TaskTpl(w http.ResponseWriter, req *http.Request) {
	var (
		param string
		sql   string
	)
	query := req.URL.Query()
	if len(query["name"]) == 0 {
		param = "1"
		sql = model.TASK_BY_ALL
	} else if strings.Contains(query["name"][0], "http") {
		param = "'%" + query["name"][0] + "%'"
		sql = model.TASK_BY_URL
	} else {
		param = "'%" + query["name"][0] + "%'"
		sql = model.TASK_BY_NAME
	}

	task := dao.GetTask(sql, param)
	views("./views/task.html", task, w)
}

func SaveTask(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for k, v := range req.Form {
		req.Form[k][0] = template.HTMLEscapeString(v[0])
	}

	taskItem := getTaskParam(req)
	if taskItem.Name == "" {
		fmt.Println("invalid params")
		return
	}

	dao.SaveTask(taskItem)
	w.Header().Add("Location", "/taskTpl")
	w.WriteHeader(302)
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
