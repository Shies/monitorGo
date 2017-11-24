package model

import (
	"fmt"
	"log"
	"strings"
)

const (
	NOTICE_BY_ALL  = "SELECT sendtype, content, tid FROM sendlist WHERE ?"
	NOITCE_BY_TID  = "SELECT sendtype, content, tid FROM sendlist WHERE tid = ?"
	_NOTICE_INSERT = "INSERT INTO sendlist(sendtype, content, tid) VALUES(?, ?, ?)"
)

type Notice struct {
	SendType int
	Content  string
	Tid      int64
}

func (d *Dao) SendList(sql string, param int64) map[int64][]*Notice {
	rows, err := d.db.Query(sql, param)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return nil
	}

	notice := make(map[int64][]*Notice)
	for rows.Next() {
		li := &Notice{}
		err := rows.Scan(&li.SendType, &li.Content, &li.Tid)
		defer rows.Close()
		if err != nil {
			fmt.Println("_return:", err.Error())
			break
		}
		notice[li.Tid] = append(notice[li.Tid], li)
	}

	return notice
}

func (d *Dao) SaveNotice(notice *Notice) bool {
	var parts []string
	if strings.Contains(notice.Content, ",") {
		parts = strings.Split(notice.Content, ",")
	}

	sql, err := d.db.Prepare(_NOTICE_INSERT)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	if len(parts) == 0 {
		for _, v := range parts {
			sql.Exec(notice.SendType, v, notice.Tid)
		}
	} else {
		sql.Exec(notice.SendType, notice.Content, notice.Tid)
	}
	return true
}
