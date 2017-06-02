package controller

import (
	"net/http"
)

func ReportTpl(w http.ResponseWriter, req *http.Request) {
	var (
		tid int64
		ip  string
	)
	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		tid = dao.GetReportTid()
	} else {
		tid = parseInt(query["tid"][0])
	}

	if len(query["ip"]) == 0 {
		ip = "0.0.0.0"
	} else {
		ip = query["ip"][0]
	}
	reports := dao.GetReportAll(tid, ip)
	views("./views/report.html", reports, w)
}

func StatusTpl(w http.ResponseWriter, req *http.Request) {
	views("./views/status.html", dao.StateReport(), w)
}

func IndexTpl(w http.ResponseWriter, req *http.Request) {
	views("./views/index.html", dao.IndexReport(), w)
}
