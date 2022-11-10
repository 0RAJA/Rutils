package work

import (
	"sync"
)

/*
	工作池
*/

type Worker struct {
	config   Config
	taskChan chan func()
	errChan  chan interface{}
	wg       sync.WaitGroup
}

type Config struct {
	TaskChanCapacity uint // 任务Chan容量
	WorkerNum        uint // 协程工人数
	ErrChanCapacity  uint // 错误Chan容量
}

func Init(config *Config) *Worker {
	w := &Worker{
		config:   *config,
		taskChan: make(chan func(), config.TaskChanCapacity),
		errChan:  make(chan interface{}, config.ErrChanCapacity),
		wg:       sync.WaitGroup{},
	}
	w.run()
	return w
}

func (w *Worker) run() {
	w.wg.Add(int(w.config.WorkerNum))
	for i := uint(0); i < w.config.WorkerNum; i++ {
		go w.work()
	}
}

func (w *Worker) SendTask(task func()) {
	w.taskChan <- task
}

func (w *Worker) work() {
	defer func() {
		err := recover()
		if err == nil {
			w.wg.Done()
			return
		}
		select {
		case w.errChan <- err:
		default:
		}
		go w.work() // 重新启动
	}()
	for task := range w.taskChan {
		task()
	}
}

func (w *Worker) Err() <-chan interface{} {
	return w.errChan
}

func (w *Worker) Stop() {
	close(w.taskChan)
	w.wg.Wait()
}

func (w *Worker) Restart(config *Config) {
	w.Stop()
	if config != nil {
		w.config = *config
	}
	w.run()
}
