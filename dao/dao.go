package dao

import (
	"database/sql"
	"monitorGo/conf"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	c 	*conf.Config
	db  *sql.DB
	err error
}

func New(c *conf.Config) *Dao {
	dao := &Dao{
		c: c,
		err: nil,
		db: nil,
	}
	dao.db, dao.err = getConn(c.Db)
	if dao.err != nil {
		panic(dao.err)
	}

	return dao
}

func getConn(db *conf.Db) (*sql.DB, error) {
	connStr := createConnStr(db.Username, db.Password, db.Addr, db.Port, db.Db_name)
	return sql.Open("mysql", connStr)
}

func createConnStr(username string, password string, addr string, port string, db_name string) string {
	return username + ":" + password + "@tcp(" + addr + ":" + port + ")/" + db_name + "?parseTime=true"
}

func convertTime(t time.Time) int64 {
	// tm2, _ := time.Parse("2006-01-02 15:04:05", t.String())
	// return tm2.Unix()
	return t.Unix()
}

