package worker

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

type mockWorkFunc struct {
	now int64
}

func (m *mockWorkFunc) DoWork() int64 {
	return m.now
}

func TestWorker_BuildWorker(t *testing.T) {
	total := 1
	w := &worker{
		url:           "http://www.example.com",
		totalReqNum:   total,
		concurrentNum: 1,
		jobsCh:        make(chan struct{}, total),
		resultCh:      make(chan int64, total),
	}
	now := time.Now().Unix()

	w.BuildWorker(&mockWorkFunc{now: now})

	w.jobsCh <- struct{}{}

	res := <-w.resultCh
	assert.Equal(t, res, now)
}

func TestWorker_BuildJobs(t *testing.T) {
	total := 5
	w := &worker{
		url:           "http://www.example.com",
		totalReqNum:   total,
		concurrentNum: 1,
		jobsCh:        make(chan struct{}, total),
		resultCh:      make(chan int64, total),
	}

	w.BuildJobs()

	var count int

	for range w.jobsCh {
		count++
		if count == total {
			break
		}
	}

	assert.Equal(t, count, total)
}

func TestWorker_PrintStatistic(t *testing.T) {
	total := 5
	w := &worker{
		url:           "http://www.example.com",
		totalReqNum:   total,
		concurrentNum: 1,
		jobsCh:        make(chan struct{}, total),
		resultCh:      make(chan int64, total),
	}

	now := time.Now().Unix()
	w.BuildWorker(&mockWorkFunc{now: now})
	w.BuildJobs()
	w.PrintStatistic()
}
