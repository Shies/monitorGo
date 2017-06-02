package model

import (
	"fmt"
	"log"
)

const (
	_CONF_BY_ALL = "SELECT setting_item, setting_value FROM settings"
	_CONF_UPDATE = "UPDATE settings SET setting_value = ? WHERE setting_item = ?"
)

type Setting struct {
	SettingItem  string
	SettingValue string
	Config       *DBField
}

type DBField struct {
	FaultTask    string
	GlobalEmails string
	GlobalPhones string
	LogDir       string
	SmsPwd       string
	SmsServer    string
	SmsUser      string
	SmtpPwd      string
	SmtpServer   string
	SmtpUser     string
}

func (d *Dao) GetConf() map[string]string {
	rows, err := d.db.Query(_CONF_BY_ALL)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return nil
	}

	resp := make(map[string]string)
	for rows.Next() {
		defer rows.Close()
		li := &Setting{}
		err := rows.Scan(&li.SettingItem, &li.SettingValue)
		if err != nil {
			fmt.Println("_return:", err.Error())
			break
		}
		resp[li.SettingItem] = li.SettingValue
	}

	return resp
}


func (d *Dao) SaveConf(key string, val string) {
	sql, err := d.db.Prepare(_CONF_UPDATE)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	sql.Exec(val, key)
}
