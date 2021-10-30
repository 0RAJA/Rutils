package RMySQL

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"                            //驱动名-这个名字其实就是数据库驱动注册到database/sql 时所使用的名字.
	userName   = "root"                             //用户名
	passWord   = "WW876001"                         //密码
	ip         = "127.0.0.1"                        //ip地址
	port       = "3306"                             //端口号
	dbName     = "cinema"                           //数据库名
	ipFormat   = "%s:%s@tcp(%s:%s)/%s?charset=utf8" //格式
)

var DB *sql.DB

func init() {
	var err error = nil
	DB, err = sql.Open(driverName, fmt.Sprintf(ipFormat, userName, passWord, ip, port, dbName))
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(100) //SetMaxOpenConns设置与数据库建立连接的最大数目。
	DB.SetMaxIdleConns(10)  //SetMaxIdleConns设置连接池中的最大闲置连接数。
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}
