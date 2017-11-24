package model

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	TASK_BY_URL  = "SELECT * FROM task WHERE status = 1 AND url LIKE ?"
	TASK_BY_NAME = "SELECT * FROM task WHERE status = 1 AND name LIKE ?"
	TASK_BY_ALL  = "SELECT * FROM task WHERE status = ?"
	_TASK_INSERT = "INSERT INTO task(name, protocol, url) VALUES(?, ?, ?)"
)

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

func (d *Dao) TaskList(sql string, param string) (tasks []*TaskItem) {
	rows, err := d.db.Query(sql, param)
	if err != nil {
		fmt.Println("DB query failed", err.Error())
		return
	}

	for rows.Next() {
		defer rows.Close()
		li := &TaskItem{}
		err := rows.Scan(&li.Id, &li.Name, &li.Protocol, &li.Url, &li.Username, &li.Password, &li.Method, &li.Params, &li.Frequency, &li.Retry, &li.Goodcode, &li.Sizerange, &li.Status, &li.Createtime, &li.Uid, &li.Gid, &li.Respbody)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, li)
	}

	return tasks
}

func (d *Dao) SaveTask(taskItem *TaskItem) {
	sql, err := d.db.Prepare(_TASK_INSERT)
	if err != nil {
		fmt.Println("invalid sql")
		return
	}

	defer sql.Close()
	sql.Exec(taskItem.Name, taskItem.Protocol, taskItem.Url)
}
