package dao

import (
	"database/sql"
	"monitorGo/conf"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	db  *sql.DB
	err error
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

func convertTime(t string) int64 {
	tm2, _ := time.Parse("2006-01-02 15:04:05", t)
	return tm2.Unix()
}
