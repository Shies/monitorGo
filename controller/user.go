package controller

import (
	"time"
	"net/http"
	"html/template"

	"monitorGo/model"
)

func userList(c Context) {
	user, _ := srv.UserList()
	group, _ := srv.GroupList()
	res := map[string]interface{}{
		"User": user,
		"Group": group,
	}
	c.SetData(res)
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
		panic("invalid params")
	}
	srv.SaveUser(user)
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
	user.LastLogin = time.Now()

	return user
}
