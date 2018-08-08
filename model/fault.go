package model

type Fault struct {
	Id            int64
	StartTime     string
	LastCheckTime string
	RespCode      int
	OutOfSize     int
	Tid           int64
	IP            string
	IsRemind      int
}
