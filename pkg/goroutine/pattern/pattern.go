package pattern

import (
	"context"
	"sync"
)

// Or 监听多个ctx 只要有一个返回消息就返回
func Or(ctx ...context.Context) context.Context {
	switch len(ctx) {
	case 0:
		return nil
	case 1:
		return ctx[0]
	}
	orCtx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		switch len(ctx) {
		case 2:
			select {
			case <-ctx[0].Done():
			case <-ctx[1].Done():
			}
		default:
			select {
			case <-ctx[0].Done():
			case <-ctx[1].Done():
			case <-ctx[2].Done():
			case <-Or(append(ctx[3:], orCtx)...).Done(): // 递归退出
			}
		}
	}()
	return orCtx
}

// Bridge 通过接受传输chan的chan，将值传递给给回去(这个是按顺序读完一个channel才会选择下一个channel)
func Bridge(ctx context.Context, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			var stream <-chan interface{}
			select {
			case mybeStream, ok := <-chanStream: // 读取chanStream中的channel
				if !ok {
					return
				}
				stream = mybeStream
			case <-ctx.Done():
				return
			}
			for val := range OrDone(ctx, stream) { // 读取channel内容发送回去
				select {
				case <-ctx.Done():
					return
				case valStream <- val:
				}
			}
		}
	}()
	return valStream
}

// Tee 读取in数据并同时发送给两个接受的channel
func Tee(ctx context.Context, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range OrDone(ctx, in) {
			var out1, out2 = out1, out2 // 本地版本，隐藏外界变量
			for i := 0; i < 2; i++ {    // 为了确保两个channel都可以被写入我们使用两次写入
				select {
				case <-ctx.Done():
					return
				case out1 <- v:
					out1 = nil // 同时写入后关闭副本channel来阻塞防止二次写入
				case out2 <- v:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

// OrDone 安全地读取c
func OrDone(ctx context.Context, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select { // 可以进行优化
				case valStream <- v:
				case <-ctx.Done():
				}
			}
		}
	}()
	return valStream
}

// FanIn 从多个channels中合并数据到一个channel
func FanIn(ctx context.Context, channels []<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

// Take 取出num个数后结束
func Take(ctx context.Context, valueStream <-chan interface{}, num int) <-chan interface{} {
	results := make(chan interface{})
	go func() {
		defer close(results)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case results <- <-valueStream:
			}
		}
	}()
	return results
}

// RepeatFn 重复调用函数
func RepeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	results := make(chan interface{})
	go func() {
		defer close(results)
		for {
			select {
			case <-ctx.Done():
				return
			case results <- fn():
			}
		}
	}()
	return results
}

// Repeat 重复生成值
func Repeat(ctx context.Context, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}
