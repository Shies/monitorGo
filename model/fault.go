package model

import "time"

type Fault struct {
	Id            int64
	StartTime     time.Time
	LastCheckTime time.Time
	RespCode      int
	OutOfSize     int
	Tid           int64
	IP            string
	IsRemind      int
}
