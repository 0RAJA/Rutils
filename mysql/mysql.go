package mysql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func mysqlInit(driverName, dataSourceName string) *sql.DB {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return conn
}

func IsNil(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
