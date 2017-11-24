package controller

import (
	"fmt"
	"html/template"
	"monitorGo/model"
	"net/http"
	"strconv"
	"time"
)

func userList(w http.ResponseWriter, req *http.Request) {
	var (
	    resp = make(map[string]interface{}, 2)
	)
	resp["User"] = dao.UserList()
	resp["Group"] = dao.GroupList()

	views("views/user.html", resp, w)
}

func saveUser(w http.ResponseWriter, req *http.Request) {
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
	w.Header().Add("Location", "/userList")
	w.WriteHeader(302)
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
