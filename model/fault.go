package model

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	_FAULT_BY_TID = "SELECT * FROM fault WHERE tid = ?"
	_FAULT_BY_ONE = "SELECT tid FROM fault WHERE 1 ORDER BY id DESC LIMIT 1"
)

type Fault struct {
	Id            int64
	StartTime     string
	LastCheckTime string
	RespCode      int
	OutOfSize     int
	Tid           int64
	IP            string
	IsRemind      int
}

func (d *Dao) GetFaultAll(tid int64, ip string) (faults []*Fault) {
	var (
		query string
		rows  *sql.Rows
	)
	if ip != "0.0.0.0" {
		query = _FAULT_BY_TID + " AND ip = ? "
		rows, _ = d.db.Query(query, tid, ip)
	} else {
		query = _FAULT_BY_TID
		rows, _ = d.db.Query(query, tid)
	}

	for rows.Next() {
		defer rows.Close()
		li := &Fault{}
		err := rows.Scan(&li.Id, &li.StartTime, &li.LastCheckTime, &li.RespCode, &li.OutOfSize, &li.Tid, &li.IP, &li.IsRemind)
		if err != nil {
			log.Fatal(err)
		}
		faults = append(faults, li)
	}

	return faults
}

func (d *Dao) GetFaultTid() (tid int64) {
	err := d.db.QueryRow(_FAULT_BY_ONE).Scan(&tid)
	if err != nil {
		fmt.Println("DB query failed:", err.Error())
		return 0
	}

	return tid
}
