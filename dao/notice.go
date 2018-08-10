package dao

import (
	"fmt"
	"log"
	"strings"
	xsql "database/sql"

	"monitorGo/model"
)

const (
	NOTICE_BY_ALL  = "SELECT sendtype, content, tid FROM sendlist WHERE ?"
	NOITCE_BY_TID  = "SELECT sendtype, content, tid FROM sendlist WHERE tid = ?"
	_NOTICE_INSERT = "INSERT INTO sendlist(`sendtype`, `content`, `tid`) VALUES(?, ?, ?)"
)

func (d *Dao) SendList(sql string, param int64) (send map[int64][]*model.Notice, err error) {
	var rows *xsql.Rows
	rows, err = d.db.Query(sql, param)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return
	}

	defer rows.Close()
	send = make(map[int64][]*model.Notice)
	for rows.Next() {
		li := &model.Notice{}
		err = rows.Scan(&li.SendType, &li.Content, &li.Tid)
		if err != nil {
			fmt.Println("_return:", err.Error())
			return
		}
		send[li.Tid] = append(send[li.Tid], li)
	}
	return
}

func (d *Dao) SaveNotice(notice *model.Notice) (err error) {
	var parts []string
	if strings.Contains(notice.Content, ",") {
		parts = strings.Split(notice.Content, ",")
	}

	sql, err := d.db.Prepare(_NOTICE_INSERT)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer sql.Close()
	if len(parts) == 0 {
		for _, v := range parts {
			sql.Exec(notice.SendType, v, notice.Tid)
		}
	} else {
		sql.Exec(notice.SendType, notice.Content, notice.Tid)
	}
	return
}
