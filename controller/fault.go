package controller

func faultList(c Context) {
	var (
		ip  string
		tid int64
		req = c.Request()
	)
	query := req.URL.Query()
	if len(query["tid"]) == 0 {
		tid = srv.FaultTid()
	} else {
		tid = parseInt(query["tid"][0])
	}
    ip = "0.0.0.0"
	if len(query["ip"]) != 0 {
	    ip = query["ip"][0]
	}
	c.SetData(srv.FaultList(tid, ip))
	c.SetPath("views/fault.html")
	views(c)
}
