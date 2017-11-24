package controller

import (
	"net/http"
)

func reportList(w http.ResponseWriter, req *http.Request) {
	var (
		tid int64
		ip  string
	)

	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		tid = dao.ReportTid()
	} else {
		tid = parseInt(query["tid"][0])
	}

	if len(query["ip"]) == 0 {
		ip = "0.0.0.0"
	} else {
		ip = query["ip"][0]
	}

	reports := dao.ReportList(tid, ip)
	views("views/report.html", reports, w)
}

func statusList(w http.ResponseWriter, req *http.Request) {
	status := dao.StateReport()
	views("views/status.html", status, w)
}

func indexList(w http.ResponseWriter, req *http.Request) {
	index := dao.IndexReport()
	views("views/index.html", index, w)
}
