package dao

import (
	"fmt"
	"log"

	"monitorGo/model"
)

const (
	_CONF_BY_ALL = "SELECT setting_item, setting_value FROM settings"
	_CONF_UPDATE = "UPDATE settings SET setting_value = ? WHERE setting_item = ?"
)

func (d *Dao) ConfList() map[string]string {
	rows, err := d.db.Query(_CONF_BY_ALL)
	if err != nil {
		fmt.Println("db query failed:", err.Error())
		return nil
	}

	resp := make(map[string]string)
	for rows.Next() {
		defer rows.Close()
		li := &model.Setting{}
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
