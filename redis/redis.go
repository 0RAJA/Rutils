package redis

import (
	"github.com/go-redis/redis"
)

type InitStruct struct {
	Host     string
	Port     string
	DB       int
	Password string
	PoolSize int
}

func Init(inits *InitStruct) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     inits.Host + ":" + inits.Port, //ip:端口
		Password: inits.Password,                //密码
		PoolSize: inits.PoolSize,                //socket最大连接数
		DB:       inits.DB,                      //默认连接数据库
	})
	_, err := rdb.Ping().Result() //测试连接
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
