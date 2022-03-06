package base

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var MysqlDB *sql.DB

func init() {
	var mysqlErr error

	MysqlDB, mysqlErr = sql.Open("mysql", "root:test2021@tcp(127.0.0.1:3306)/hero_story")

	if mysqlErr != nil {
		panic(mysqlErr)
	}

	MysqlDB.SetMaxOpenConns(128)
	MysqlDB.SetMaxIdleConns(16)
	MysqlDB.SetConnMaxLifetime(2 * time.Minute)

	if mysqlErr = MysqlDB.Ping(); mysqlErr != nil {
		panic(mysqlErr)
	}
}
