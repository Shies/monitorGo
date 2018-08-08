package dao

import (
	"fmt"
	"log"

	"monitorGo/model"
)

const (
	_GROUPS_BY_ALL = "SELECT * FROM usergroup WHERE 1"
	_GROUPS_INSERT = "INSERT INTO usergroup(name, is_user_admin, is_group_admin, is_conf_admin) VALUES(?, ?, ?, ?)"
)

func (d *Dao) GroupList() []*model.Group {
	rows, err := d.db.Query(_GROUPS_BY_ALL)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return nil
	}

	groups := []*model.Group{}
	for rows.Next() {
		defer rows.Close()
		li := &model.Group{}
		err := rows.Scan(&li.Id, &li.Name, &li.IsUserAdmin, &li.IsGroupAdmin, &li.IsConfAdmin)
		if err != nil {
			fmt.Println(err)
			break
		}
		groups = append(groups, li)
	}

	return groups
}


func (d *Dao) SaveGroup(group *model.Group) {
	sql, err := d.db.Prepare(_GROUPS_INSERT)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	sql.Exec(group.Name, group.IsUserAdmin, group.IsGroupAdmin, group.IsConfAdmin)
}