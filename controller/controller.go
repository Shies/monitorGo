package controller

import (
	_ "fmt"
	"log"
	"strconv"
	"net/http"
	"html/template"

	"monitorGo/service"
	"monitorGo/conf"
	"context"
)

var (
	srv *service.Service
)

type Context interface {
	context.Context
    Request() *http.Request
    Response() http.ResponseWriter
    SetPath(string)
    SetData(interface{})
    GetPath() string
    GetData() interface{}
}

type implContext struct {
	context.Context
    req  *http.Request
    res  http.ResponseWriter
    path string
    data interface{}
}

func (ic *implContext) Request() *http.Request {
    return ic.req
}

func (ic *implContext) Response() http.ResponseWriter {
    return ic.res
}

func (ic *implContext) SetPath(path string) {
    ic.path = path
}

func (ic *implContext) SetData(data interface{}) {
    ic.data = data
}

func (ic *implContext) GetPath() string {
	return ic.path
}

func (ic *implContext) GetData() interface{} {
	return ic.data
}

func parseInt(value string) int64 {
	intval, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		intval = 0
	}

	return intval
}

func atoi(value string) int {
	intval, err := strconv.Atoi(value)
	if err != nil {
		intval = 0
	}

	return intval
}

func views(c Context) error {
	t, err := template.ParseFiles("./" + c.GetPath())
	if err != nil {
		panic(err)
	}

	return t.Execute(c.Response(), c.GetData())
}

func initService() {
	srv = service.New(&conf.Conf)
}

func Register() bool {
	conf.ParseConfig()
	initService()
	setHttpHandle()

	// 设置监听端口
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	return true
}

func setHttpHandle() error {
    var (
        err error
        router = make(map[string]interface{})
    )
    router["userTpl"] = userList
    router["saveUser"] = saveUser
    router["groupTpl"] = groupList
    router["saveGroup"] = saveGroup
    router["confTpl"] = confList
    router["saveConf"] = saveConf
    router["noticeTpl"] = noticeList
    router["saveNotice"] = saveNotice
    router["ipTpl"] = ipList
    router["saveIP"] = saveIP
    router["taskTpl"] = taskList
    router["saveTask"] = saveTask
    router["reportTpl"] = reportList
    router["faultTpl"] = faultList
    router["statusTpl"] = statusList
    router["indexTpl"] = indexList
    for route, function := range router {
        f := function.(func(Context))
        http.HandleFunc("/" + route, func(w http.ResponseWriter, req *http.Request) {
        	f(&implContext{req: req, res: w})
        })
    }

	return err
}
