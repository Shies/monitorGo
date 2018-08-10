package dao

import (
	"fmt"
	"log"
	xsql "database/sql"

	"monitorGo/model"
)

const (
	_CONF_BY_ALL = "SELECT setting_item, setting_value FROM settings"
	_CONF_UPDATE = "UPDATE settings SET setting_value = ? WHERE setting_item = ?"
)

func (d *Dao) ConfList() (sets map[string]string, err error) {
	var rows *xsql.Rows
	rows, err = d.db.Query(_CONF_BY_ALL)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return
	}

	defer rows.Close()
	sets = make(map[string]string)
	for rows.Next() {
		li := &model.Setting{}
		err = rows.Scan(&li.SettingItem, &li.SettingValue)
		if err != nil {
			fmt.Println("_return:", err.Error())
			return
		}
		sets[li.SettingItem] = li.SettingValue
	}
	return
}

func (d *Dao) SaveConf(key string, val string) (err error) {
	sql, err := d.db.Prepare(_CONF_UPDATE)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(val, key)
	return
}
