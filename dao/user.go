package dao

import (
	"fmt"
	"log"
	xsql "database/sql"

	"monitorGo/model"
)

const (
	_USERS_BY_ALL = "SELECT * FROM user WHERE 1"
	_USERS_INSERT = "INSERT INTO user(loginname, name, email, phone, edit_group_task, edit_group_user, gid, lastlogin) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
)

func (d *Dao) UserList() (users []*model.User, err error) {
	var rows *xsql.Rows
	rows, err = d.db.Query(_USERS_BY_ALL)
	if err != nil {
		fmt.Println("DB query failed", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		li := &model.User{}
		err = rows.Scan(&li.Id, &li.LastLogin, &li.Name, &li.Email, &li.Phone, &li.EditGroupTask, &li.EditGroupUser, &li.Gid, &li.LastLogin)
		if err != nil {
			fmt.Println(err)
			return
		}
		users = append(users, li)
	}
	return
}

func (d *Dao) SaveUser(user *model.User) (err error) {
	sql, err := d.db.Prepare(_USERS_INSERT)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer sql.Close()
	_, err = sql.Exec(user.LastLogin, user.Name, user.Email, user.Phone, user.EditGroupTask, user.EditGroupUser, user.Gid, user.LastLogin)
	return
}
