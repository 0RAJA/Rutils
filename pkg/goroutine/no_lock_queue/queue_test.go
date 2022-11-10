package no_lock_queue_test

import (
	"math/rand"
	"testing"
	"time"

	queue "github.com/0RAJA/Rutils/pkg/goroutine/no_lock_queue"
	"github.com/0RAJA/Rutils/pkg/goroutine/work"
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/stretchr/testify/require"
)

func countResult(done, in, out <-chan int64) bool {
	// var inSlice, outSlice []int64
over:
	for {
		select {
		case <-done:
			break over
		case v, ok := <-in:
			if ok {

			}
		case v, ok := <-out:

		}
	}
}
func TestNewLKQueue(t *testing.T) {
	q := queue.NewLKQueue()
	nums := make([]int64, utils.RandomInt(5, 100))
	for i := range nums {
		nums[i] = utils.RandomInt(1, 100)
	}
	gNum := utils.RandomInt(5, 10)
	w := InitWorker(0, uint(gNum))
	for _, v := range nums {
		w.SendTask(func() {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
			q.Enqueue(v)
		})
	}
	time.Sleep(1 * time.Second)
	require.EqualValues(t, len(nums), q.Count())
	for range nums {
		w.SendTask(func() {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
			q.Dequeue()
		})
	}
}

func InitWorker(taskChanCapacity, workerNum uint) *work.Worker {
	config := &work.Config{
		TaskChanCapacity: taskChanCapacity,
		WorkerNum:        workerNum,
		ErrChanCapacity:  0,
	}
	return work.Init(config)
}
