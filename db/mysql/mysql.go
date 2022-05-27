package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlInit(driverName, dataSourceName string, maxOpenConns, maxIdleCOnns int) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleCOnns)
	db.SetConnMaxLifetime(time.Minute * 10)
	return db
}
