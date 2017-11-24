package model

import (
	"fmt"
	"log"
)

const (
	_USERS_BY_ALL = "SELECT * FROM user WHERE 1"
	_USERS_INSERT = "INSERT INTO user(loginname, name, email, phone, edit_group_task, edit_group_user, gid, lastlogin) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
)

type User struct {
	Id            int64
	LoginName     string
	Name          string
	Email         string
	Phone         string
	EditGroupTask int
	EditGroupUser int
	Gid           int
	LastLogin     string
}

func (d *Dao) UserList() []*User {
	rows, err := d.db.Query(_USERS_BY_ALL)
	if err != nil {
		fmt.Println("DB query failed", err.Error())
		return nil
	}

	users := []*User{}
	for rows.Next() {
		defer rows.Close()
		li := &User{}
		err = rows.Scan(&li.Id, &li.LastLogin, &li.Name, &li.Email, &li.Phone, &li.EditGroupTask, &li.EditGroupUser, &li.Gid, &li.LastLogin)
		if err != nil {
			fmt.Println(err)
			break
		}
		users = append(users, li)
	}

	return users
}

func (d *Dao) SaveUser(user *User) {
	sql, err := d.db.Prepare(_USERS_INSERT)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	sql.Exec(user.LastLogin, user.Name, user.Email, user.Phone, user.EditGroupTask, user.EditGroupUser, user.Gid, user.LastLogin)
}
