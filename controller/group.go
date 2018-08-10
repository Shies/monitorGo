package controller

import (
	"net/http"
	"html/template"

	"monitorGo/model"
)

func groupList(c Context) {
	groups, _ := srv.GroupList()
	c.SetData(groups)
	c.SetPath("views/group.html")
	views(c)
}

func saveGroup(c Context) {
	var (
		req = c.Request()
		res = c.Response()
	)
	req.ParseForm()
	group := getGroupParams(req)
	if group.Name == "" {
		panic("invalid params")
	}
	srv.SaveGroup(group)
	res.Header().Add("Location", "/groupTpl")
	res.WriteHeader(302)
}

func getGroupParams(req *http.Request) *model.Group {
	group := &model.Group{}
	group.Name = template.HTMLEscapeString(req.PostFormValue("name"))
	group.IsUserAdmin = atoi(req.PostFormValue("is_user_admin"))
	group.IsGroupAdmin = atoi(req.PostFormValue("is_group_admin"))
	group.IsConfAdmin = atoi(req.PostFormValue("is_conf_admin"))

	return group
}
