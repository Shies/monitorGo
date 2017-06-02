package model

import (
	"database/sql"
	"monitorGo/conf"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	db  *sql.DB
	err error
	req *http.Request
}

func New() *Dao {
	dao := &Dao{}
	dao.db, dao.err = getConn()
	if dao.err != nil {
		panic(dao.err)
	}

	return dao
}

func getConn() (*sql.DB, error) {
	conf.ParseConfig()
	connStr := createConnStr(conf.Conf.Db.Username, conf.Conf.Db.Password, conf.Conf.Db.Addr, conf.Conf.Db.Port, conf.Conf.Db.Db_name)
	return sql.Open("mysql", connStr)
}

func createConnStr(username string, password string, addr string, port string, db_name string) string {
	return username + ":" + password + "@tcp(" + addr + ":" + port + ")/" + db_name
}
