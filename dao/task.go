package dao

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"monitorGo/model"
)

const (
	TASK_BY_URL  = "SELECT * FROM task WHERE status = 1 AND url LIKE ?"
	TASK_BY_NAME = "SELECT * FROM task WHERE status = 1 AND name LIKE ?"
	TASK_BY_ALL  = "SELECT * FROM task WHERE status = ?"
	_TASK_INSERT = "INSERT INTO task(name, protocol, url) VALUES(?, ?, ?)"
)

func (d *Dao) TaskList(sql string, param string) (tasks []*model.TaskItem) {
	rows, err := d.db.Query(sql, param)
	if err != nil {
		fmt.Println("DB query failed", err.Error())
		return
	}

	for rows.Next() {
		defer rows.Close()
		li := &model.TaskItem{}
		err := rows.Scan(&li.Id, &li.Name, &li.Protocol, &li.Url, &li.Username, &li.Password, &li.Method, &li.Params, &li.Frequency, &li.Retry, &li.Goodcode, &li.Sizerange, &li.Status, &li.Createtime, &li.Uid, &li.Gid, &li.Respbody)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, li)
	}

	return tasks
}

func (d *Dao) SaveTask(taskItem *model.TaskItem) {
	sql, err := d.db.Prepare(_TASK_INSERT)
	if err != nil {
		fmt.Println("invalid sql")
		return
	}

	defer sql.Close()
	sql.Exec(taskItem.Name, taskItem.Protocol, taskItem.Url)
}
