package model

// task model
type TaskItem struct {
	Id         int64
	Name       string
	Protocol   string
	Url        string
	Username   string
	Password   string
	Method     string
	Params     string
	Frequency  int
	Retry      int
	Goodcode   int
	Sizerange  string
	Status     int
	Createtime string
	Uid        int
	Gid        int
	Respbody   string
}
