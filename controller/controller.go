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
	_int64, _ := strconv.ParseInt(value, 10, 64)
	return _int64
}

func atoi(value string) int {
	_int, _ := strconv.Atoi(value)
	return _int
}

func views(path string, data interface{}, w http.ResponseWriter) error {
	t, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return t.Execute(w, data)
}

func Register() bool {
	setHttpHandle()

	// 设置监听端口
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	return true
}

func setHttpHandle() bool {
	http.HandleFunc("/userTpl", UserTpl)
	http.HandleFunc("/saveUser", SaveUser)
	http.HandleFunc("/groupTpl", GroupTpl)
	http.HandleFunc("/saveGroup", saveGroup)
	http.HandleFunc("/confTpl", ConfTpl)
	http.HandleFunc("/saveConf", SaveConf)
	http.HandleFunc("/noticeTpl", NoticeTpl)
	http.HandleFunc("/saveNotice", SaveNotice)
	http.HandleFunc("/ipTpl", IPTpl)
	http.HandleFunc("/saveIP", SaveIP)
	http.HandleFunc("/taskTpl", TaskTpl)
	http.HandleFunc("/saveTask", SaveTask)
	http.HandleFunc("/reportTpl", ReportTpl)
	http.HandleFunc("/faultTpl", FaultTpl)
	http.HandleFunc("/statusTpl", StatusTpl)
	http.HandleFunc("/indexTpl", IndexTpl)
	return true
}
