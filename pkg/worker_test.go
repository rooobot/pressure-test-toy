package pkg

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
	w := &Worker{
		URL:           "http://www.example.com",
		TotalReqNum:   total,
		ConcurrentNum: 1,
		JobsCh:        make(chan struct{}, total),
		ResultCh:      make(chan int64, total),
	}
	now := time.Now().Unix()

	w.BuildWorker(&mockWorkFunc{now: now})

	w.JobsCh <- struct{}{}

	res := <-w.ResultCh
	assert.Equal(t, res, now)
}

func TestWorker_BuildJobs(t *testing.T) {
	total := 5
	w := &Worker{
		URL:           "http://www.example.com",
		TotalReqNum:   total,
		ConcurrentNum: 1,
		JobsCh:        make(chan struct{}, total),
		ResultCh:      make(chan int64, total),
	}

	//w.BuildWorker(&mockWorkFunc{now: time.Now().Unix()})

	w.BuildJobs()

	var count int

	for range w.JobsCh {
		count++
		if count == total {
			break
		}
	}

	assert.Equal(t, count, total)
}

func TestWorker_PrintStatistic(t *testing.T) {
	total := 5
	w := &Worker{
		URL:           "http://www.example.com",
		TotalReqNum:   total,
		ConcurrentNum: 1,
		JobsCh:        make(chan struct{}, total),
		ResultCh:      make(chan int64, total),
	}

	now := time.Now().Unix()
	w.BuildWorker(&mockWorkFunc{now: now})
	w.BuildJobs()
	w.PrintStatistic()
}
