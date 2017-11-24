package controller

import (
	"fmt"
	"html/template"
	"log"
	"monitorGo/model"
	"net/http"
	"strconv"
)

var (
	dao = model.New()
)

func parseInt(value string) int64 {
	intval, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		intval = 0
	}

	return intval
}

func atoi(value string) (intval int) {
	intval, err := strconv.Atoi(value)
	if err != nil {
		intval = 0
	}

	return intval
}

func views(path string, data interface{}, w http.ResponseWriter) error {
	t, err := template.ParseFiles("./" + path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return t.Execute(w, data)
}

func Register() bool {
	setHttpHandle()

	// 设置监听端口
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	return true
}

func setHttpHandle() {
	http.HandleFunc("/userTpl", userList)
	http.HandleFunc("/saveUser", saveUser)
	http.HandleFunc("/groupTpl", groupList)
	http.HandleFunc("/saveGroup", saveGroup)
	http.HandleFunc("/confTpl", confList)
	http.HandleFunc("/saveConf", saveConf)
	http.HandleFunc("/noteTpl", noteList)
	http.HandleFunc("/saveNote", saveNote)
	http.HandleFunc("/ipTpl", ipList)
	http.HandleFunc("/saveIP", saveIP)
	http.HandleFunc("/taskTpl", taskList)
	http.HandleFunc("/saveTask", saveTask)
	http.HandleFunc("/reportTpl", reportList)
	http.HandleFunc("/faultTpl", faultList)
	http.HandleFunc("/statusTpl", statusList)
	http.HandleFunc("/indexTpl", indexList)
	return
}
