package model

import "time"

type Report struct {
	Id       int64
	Time     time.Time
	RespTime float64
	RespCode int
	Size     int
	Tid      int64
	IP       string
}

type Status struct {
	RespTime float64
	RespCode int
	IP       string
	Name     string
	Url      string
	GoodCode int
}

type Index struct {
	Id        int64
	Name      string
	Url       string
	Avail     string
	TotalTime string
}