package controller

import (
	"fmt"
	"html/template"
	"monitorGo/model"
	"net/http"
)

func GroupTpl(w http.ResponseWriter, req *http.Request) {
	group := dao.GetGroup()
	views("./views/group.html", group, w)
}

func saveGroup(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	group := getGroupParams(req)
	if group.Name == "" {
		fmt.Println("invalid params")
		return
	}

	dao.SaveGroup(group)
	w.Header().Add("Location", "/groupTpl")
	w.WriteHeader(302)
}

func getGroupParams(req *http.Request) *model.Group {
	group := &model.Group{}
	group.Name = template.HTMLEscapeString(req.PostFormValue("name"))
	group.IsUserAdmin = atoi(req.PostFormValue("is_user_admin"))
	group.IsGroupAdmin = atoi(req.PostFormValue("is_group_admin"))
	group.IsConfAdmin = atoi(req.PostFormValue("is_conf_admin"))

	return group
}
