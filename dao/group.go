package dao

import (
	"fmt"
	"log"
	xsql "database/sql"

	"monitorGo/model"
)

const (
	_GROUPS_BY_ALL = "SELECT * FROM usergroup WHERE 1"
	_GROUPS_INSERT = "INSERT INTO usergroup(name, is_user_admin, is_group_admin, is_conf_admin) VALUES(?, ?, ?, ?)"
)

func (d *Dao) GroupList() (groups []*model.Group, err error) {
	var rows *xsql.Rows
	rows, err = d.db.Query(_GROUPS_BY_ALL)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		li := &model.Group{}
		err = rows.Scan(&li.Id, &li.Name, &li.IsUserAdmin, &li.IsGroupAdmin, &li.IsConfAdmin)
		if err != nil {
			fmt.Println(err)
			return
		}
		groups = append(groups, li)
	}
	return
}

func (d *Dao) SaveGroup(group *model.Group) (err error) {
	sql, err := d.db.Prepare(_GROUPS_INSERT)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(group.Name, group.IsUserAdmin, group.IsGroupAdmin, group.IsConfAdmin)
	return
}