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

var rdb *redis.Client

func Init(inits *InitStruct) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     inits.Host + ":" + inits.Port, //ip:端口
		Password: inits.Password,                //密码
		PoolSize: inits.PoolSize,                //连接池
		DB:       inits.DB,                      //默认连接数据库
	})
	_, err := rdb.Ping().Result() //测试连接
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}

func GetPipe() redis.Pipeliner {
	return rdb.TxPipeline()
}
