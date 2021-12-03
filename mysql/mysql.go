package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// IpFormat parseTime=true 不加导致时间戳无法转换
const (
	IpFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=%v&charset=%s" //格式
)

type InitStruct struct {
	DriverName   string //驱动名
	Username     string //填写你的数据库账号
	Password     string // 填写你的数据库密码
	Host         string
	Port         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

var db *sqlx.DB

func Init(initSetting *InitStruct) error {
	var err error
	db, err = sqlx.Connect(initSetting.DriverName, fmt.Sprintf(IpFormat, initSetting.Username, initSetting.Password, initSetting.Host, initSetting.Port, initSetting.DBName, initSetting.ParseTime, initSetting.Charset))
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(initSetting.MaxOpenConns) //SetMaxOpenConns设置与数据库建立连接的最大数目。
	db.SetMaxIdleConns(initSetting.MaxIdleConns) //SetMaxIdleConns设置连接池中的最大闲置连接数。
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = db.Close()
}
