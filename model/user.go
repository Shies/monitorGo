package model

import "time"

type User struct {
	Id            int64
	LoginName     string
	Name          string
	Email         string
	Phone         string
	EditGroupTask int
	EditGroupUser int
	Gid           int
	LastLogin     time.Time
}
