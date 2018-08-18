package dao

import (
	"fmt"
	"log"
	xsql "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"monitorGo/model"
)

const (
	TASK_BY_URL  = "SELECT * FROM task WHERE status = 1 AND url LIKE ?"
	TASK_BY_NAME = "SELECT * FROM task WHERE status = 1 AND name LIKE ?"
	TASK_BY_ALL  = "SELECT * FROM task WHERE status = ?"
	_TASK_INSERT = "INSERT INTO task(name, protocol, url) VALUES(?, ?, ?)"
)

func (d *Dao) TaskList(sql string, param string) (tasks []*model.TaskItem, err error) {
	var rows *xsql.Rows
	rows, err = d.db.Query(sql, param)
	if err != nil {
		fmt.Println("DB query failed", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		li := &model.TaskItem{}
		err = rows.Scan(&li.Id, &li.Name, &li.Protocol, &li.Url, &li.Username, &li.Password, &li.Method, &li.Params, &li.Frequency, &li.Retry, &li.Goodcode, &li.Sizerange, &li.Status, &li.Createtime, &li.Uid, &li.Gid, &li.Respbody)
		if err != nil {
			log.Fatal(err)
			return
		}
		tasks = append(tasks, li)
	}
	return
}

func (d *Dao) TaskTick(sql string, param string) (tasks map[int64]*model.TaskItem, err error) {
	var rows *xsql.Rows
	rows, err = d.db.Query(sql, param)
	if err != nil {
		fmt.Println("DB query failed", err.Error())
		return
	}

	defer rows.Close()
	tasks = make(map[int64]*model.TaskItem)
	for rows.Next() {
		li := &model.TaskItem{}
		err = rows.Scan(&li.Id, &li.Name, &li.Protocol, &li.Url, &li.Username, &li.Password, &li.Method, &li.Params, &li.Frequency, &li.Retry, &li.Goodcode, &li.Sizerange, &li.Status, &li.Createtime, &li.Uid, &li.Gid, &li.Respbody)
		if err != nil {
			log.Fatal(err)
			return
		}
		tasks[li.Id] = li
	}
	return
}

func (d *Dao) SaveTask(taskItem *model.TaskItem) (err error) {
	sql, err := d.db.Prepare(_TASK_INSERT)
	if err != nil {
		fmt.Println("invalid sql")
		return
	}

	defer sql.Close()
	_, err = sql.Exec(taskItem.Name, taskItem.Protocol, taskItem.Url)
	return
}
