package limit_test

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/0RAJA/Rutils/RSync/limit"
	"golang.org/x/time/rate"
)

// 通用api
type Api interface {
	ReadFile(ctx context.Context) error       // 模拟读文件
	ResolveAddress(ctx context.Context) error // 模拟网络请求
}

// 测试api实例
type testApi struct {
	netWorkLimit, diskLimit, apiLimit limit.RateLimiter // 多个维度进行限流
}

func Open() Api {
	apiLimit := limit.MultiLimiter(
		rate.NewLimiter(limit.Per(2, time.Second), 1),   // 每秒的限制,防止突发请求,每1秒补充两个
		rate.NewLimiter(limit.Per(10, time.Minute), 10), // 每分钟的限制，设置初始池,每10秒补充一个
	)
	diskLimit := limit.MultiLimiter(
		rate.NewLimiter(rate.Limit(1), 1),
	)
	netWorkLimit := limit.MultiLimiter(
		rate.NewLimiter(limit.Per(3, time.Second), 3),
	)
	return &testApi{
		apiLimit:     apiLimit,
		diskLimit:    diskLimit,
		netWorkLimit: netWorkLimit,
	}
}

func (t *testApi) ReadFile(ctx context.Context) error {
	if err := limit.MultiLimiter(t.apiLimit, t.diskLimit).Wait(ctx); err != nil { // 融合api限流和磁盘限流
		return err
	}
	return nil
}

func (t *testApi) ResolveAddress(ctx context.Context) error {
	if err := limit.MultiLimiter(t.apiLimit, t.netWorkLimit).Wait(ctx); err != nil {
		return err
	}
	return nil
}

// 模拟20个并发api请求，其中10个用于读取文件，10个用于网络请求
func ExampleMultiLimiter() {
	defer log.Println("Done")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.Lshortfile)
	apiConn := Open()
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConn.ReadFile(context.Background()); err != nil {
				log.Println("cannot read file:", err)
				return
			}
			log.Println("read file")
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConn.ResolveAddress(context.Background()); err != nil {
				log.Println("cannot resolve address:", err)
				return
			}
			log.Println("ResolveAddress")
		}()
	}
	wg.Wait()
}
