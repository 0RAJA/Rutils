package mysql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func mysqlInit(driverName, dataSourceName string, maxOpenConns, maxIdleCOnns int) *sql.DB {
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

func IsNil(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
