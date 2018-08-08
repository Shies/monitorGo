package dao

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"monitorGo/model"
)

const (
	IPS_BY_ALL       = "SELECT tid, ip FROM taskip WHERE ?"
	IPS_BY_TID       = "SELECT tid, ip FROM taskip WHERE tid = ?"
	_IPS_INSERT      = "INSERT INTO taskip(tid, ip) VALUES(?, ?)"
	IPS_COUNT_BY_TID = "SELECT COUNT(*) AS total FROM taskip WHERE tid = ?"
)

func (d *Dao) TaskIP(query string, param int64) map[int64][]*model.TaskIP {
	rows, err := d.db.Query(query, param)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return nil
	}

	var ips = make(map[int64][]*model.TaskIP)
	for rows.Next() {
		defer rows.Close()
		li := &model.TaskIP{}
		err := rows.Scan(&li.Tid, &li.IP)
		if err != nil {
			fmt.Println("Scan failed:", err.Error())
			break
		}
		ips[li.Tid] = append(ips[li.Tid], li)
	}

	return ips
}

func (d *Dao) SaveIP(IP *model.TaskIP) {
	var parts []string
	if strings.Contains(IP.IP, ",") {
		parts = strings.Split(IP.IP, ",")
	}

	sql, err := d.db.Prepare(_IPS_INSERT)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	if len(parts) > 0 {
		for _, v := range parts {
			sql.Exec(IP.Tid, v)
		}
	} else {
		sql.Exec(IP.Tid, IP.IP)
	}
}
