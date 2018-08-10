package dao

import (
	"fmt"
	"log"
	"database/sql"

	"monitorGo/model"
)

const (
	_FAULT_BY_TID = "SELECT * FROM fault WHERE tid = ?"
	_FAULT_BY_ONE = "SELECT tid FROM fault WHERE 1 ORDER BY id DESC LIMIT 1"
)

func (d *Dao) FaultList(tid int64, ip string) (faults []*model.Fault, err error) {
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

	defer rows.Close()
	for rows.Next() {
		li := &model.Fault{}
		err = rows.Scan(&li.Id, &li.StartTime, &li.LastCheckTime, &li.RespCode, &li.OutOfSize, &li.Tid, &li.IP, &li.IsRemind)
		if err != nil {
			log.Fatal(err)
			return
		}
		faults = append(faults, li)
	}
	return
}

func (d *Dao) FaultTid() (tid int64, err error) {
	err = d.db.QueryRow(_FAULT_BY_ONE).Scan(&tid)
	if err != nil {
		fmt.Println("DB query failed:", err.Error())
		return 0, err
	}
	return
}
