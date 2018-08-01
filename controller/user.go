package controller

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"html/template"

	"monitorGo/model"
)

func userList(c Context) {
	resp := make(map[string]interface{})
	resp["User"] = dao.GetUser()
	resp["Group"] = dao.GetGroup()
	c.SetData(resp)
	c.SetPath("views/user.html")
	views(c)
}

func saveUser(c Context) {
	var (
		req = c.Request()
		res = c.Response()
	)
	req.ParseForm()
	for k, v := range req.Form {
		req.Form[k][0] = template.HTMLEscapeString(v[0])
	}
	user := getUserParams(req)
	if user.Name == "" || user.LoginName == "" {
		fmt.Println("invalid params")
		return
	}
	dao.SaveUser(user)
	res.Header().Add("Location", "/userTpl")
	res.WriteHeader(302)
}

func getUserParams(req *http.Request) (user *model.User) {
	user = &model.User{}
	user.LoginName = req.PostFormValue("loginname")
	user.Name = req.PostFormValue("name")
	user.Email = req.PostFormValue("email")
	user.Phone = req.PostFormValue("phone")
	user.EditGroupTask = atoi(req.PostFormValue("edit_group_task"))
	user.EditGroupUser = atoi(req.PostFormValue("edit_group_user"))
	user.Gid = atoi(req.PostFormValue("gid"))
	user.LastLogin = strconv.FormatInt(time.Now().UnixNano(), 10)

	return user
}
