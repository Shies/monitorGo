package controller

import "net/http"

func faultList(w http.ResponseWriter, req *http.Request) {
	var (
		tid int64
		ip  string
	)

	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		tid = dao.FaultTid()
	} else {
		tid = parseInt(query["tid"][0])
	}

	if len(query["ip"]) != 0 {
		ip = query["ip"][0]
	} else {
		ip = "0.0.0.0"
	}

	faults := dao.FaultList(tid, ip)
	views("views/fault.html", faults, w)
}
