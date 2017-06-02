package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_REPORT_BY_ONE     = "SELECT tid FROM report WHERE 1 ORDER BY id DESC LIMIT 1"
	_REPORT_BY_TID     = "SELECT * FROM report WHERE tid = ?"
	_REPORT_BY_ONE_TID = "SELECT resptime, respcode, ip FROM report WHERE tid = ? ORDER BY id DESC LIMIT 1"
	_REPORT_BY_ALL_TID = "SELECT * FROM report WHERE tid = ? ORDER BY id DESC LIMIT ?"
)

type Report struct {
	Id       int64
	Time     string
	RespTime float64
	RespCode int
	Size     int
	Tid      int64
	IP       string
}

type Status struct {
	RespTime float64
	RespCode int
	IP       string
	Name     string
	Url      string
	GoodCode int
}

type Index struct {
	Id        int64
	Name      string
	Url       string
	Avail     string
	TotalTime string
}

func (d *Dao) GetReportAll(tid int64, ip string) (reports []*Report) {
	var (
		query string
		rows  *sql.Rows
	)
	if ip != "0.0.0.0" {
		query = _REPORT_BY_TID + " AND ip = ? "
		rows, _ = d.db.Query(query, tid, ip)
	} else {
		query = _REPORT_BY_TID
		rows, _ = d.db.Query(query, tid)
	}

	for rows.Next() {
		defer rows.Close()
		li := &Report{}
		err := rows.Scan(&li.Id, &li.Time, &li.RespTime, &li.RespCode, &li.Size, &li.Tid, &li.IP)
		if err != nil {
			log.Fatal(err)
		}
		reports = append(reports, li)
	}

	return reports
}

func (d *Dao) GetReportTid() (tid int64) {
	row := d.db.QueryRow(_REPORT_BY_ONE)
	if err := row.Scan(&tid); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
		} else {
			fmt.Println("DB query failed:", err.Error())
		}
		return
	}

	return tid
}

func (d *Dao) getIPCountByTid(tid int64) (count int64) {
	row := d.db.QueryRow(IPS_COUNT_BY_TID, tid)
	if err := row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
		} else {
			fmt.Println("DB query failed:", err.Error())
		}
		return
	}

	return count
}

func (d *Dao) getReportOneByTid(tid int64) (r *Report) {
	r = &Report{}
	row := d.db.QueryRow(_REPORT_BY_ONE_TID, tid)
	if err := row.Scan(&r.RespTime, &r.RespCode, &r.IP); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("now rows")
		} else {
			fmt.Println("DB query failed:", err.Error())
		}
		return
	}

	return r
}

func (d *Dao) getReportAllByTid(tid int64, count int64) []*Report {
	rows, err := d.db.Query(_REPORT_BY_ALL_TID, tid, count)
	if err != nil {
		fmt.Println("DB query failed:", err.Error())
		return nil
	}

	var reports []*Report
	for rows.Next() {
		defer rows.Close()
		li := &Report{}
		err := rows.Scan(&li.Id, &li.Time, &li.RespTime, &li.RespCode, &li.Size, &li.Tid, &li.IP)
		if err != nil {
			log.Fatal(err)
		}
		reports = append(reports, li)
	}

	return reports
}

func (d *Dao) Assign(ti *TaskItem, r *Report) (li *Status) {
	li = &Status{}
	li.RespTime = r.RespTime
	li.RespCode = r.RespCode
	li.IP = r.IP
	li.Name = ti.Name
	li.Url = ti.Url
	li.GoodCode = ti.Goodcode

	return li
}

func (d *Dao) StateReport() []*Status {
	tasks := d.GetTask(TASK_BY_ALL, "1")
	if tasks == nil {
		fmt.Println("tasks for nil")
		return nil
	}

	var status []*Status
	for _, v := range tasks {
		var li *Status
		count := d.getIPCountByTid(v.Id)
		if count == 0 {
			r := d.getReportOneByTid(v.Id)
			li = d.Assign(v, r)
			status = append(status, li)
		} else {
			all := d.getReportAllByTid(v.Id, count)
			for _, r := range all {
				li := d.Assign(v, r)
				status = append(status, li)
			}
		}
	}

	return status
}

func (d *Dao) convertTime(t string) int64 {
	tm2, _ := time.Parse("2006-01-02 15:04:05", t)
	return tm2.Unix()
}

func (d *Dao) IndexAssign(v *TaskItem, now time.Time) (li *Index) {
	li = &Index{}
	li.Id = v.Id
	li.Name = v.Name
	li.Url = v.Url

	totaltime := now.Unix() - d.convertTime(v.Createtime)
	ip_count := d.getIPCountByTid(v.Id)
	if ip_count == 0 {
		ip_count = 1
	}

	var faulttime int64
	faults := d.GetFaultAll(v.Id, "0.0.0.0")
	for _, val := range faults {
		timediff := d.convertTime(val.LastCheckTime) - d.convertTime(val.StartTime) + int64(60*(v.Frequency))
		faulttime = int64(faulttime) + int64(timediff)
		if faulttime == 0 {
			continue
		}
	}

	li.Avail = fmt.Sprintf("%.3f", float64(float64(1)-float64(faulttime)/float64(totaltime)/float64(ip_count))*100)
	li.TotalTime = fmt.Sprintf("%.3f", float64(float64(totaltime)/float64(60)))

	return li
}

func (d *Dao) IndexReport() []*Index {
	tasks := d.GetTask(TASK_BY_ALL, "1")
	if tasks == nil {
		fmt.Println("tasks for nil")
		return nil
	}

	var index []*Index
	now := time.Now()
	for _, v := range tasks {
		li := d.IndexAssign(v, now)
		index = append(index, li)
	}

	return index
}
