package controller

func reportList(c Context) {
	var (
		tid int64
		ip  string
		req = c.Request()
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
	c.SetData(reports)
	c.SetPath("views/report.html")
	views(c)
}

func statusList(c Context) {
	c.SetData(dao.StateReport())
	c.SetPath("views/status.html")
	views(c)
}

func indexList(c Context) {
	c.SetData(dao.IndexReport())
	c.SetPath("views/index.html")
	views(c)
}
