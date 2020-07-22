package worker

import "fmt"

type worker struct {
	url           string
	totalReqNum   int
	concurrentNum int
	jobsCh        chan struct{}
	resultCh      chan int64
}

func NewWorker(url string, concurrentNum int, totalReqNum int) *worker {
	return &worker{
		url:           url,
		concurrentNum: concurrentNum,
		totalReqNum:   totalReqNum,
		jobsCh:        make(chan struct{}, totalReqNum),
		resultCh:      make(chan int64, totalReqNum),
	}
}

type WorkFunc interface {
	DoWork() int64
}

func (w *worker) BuildWorker(wf WorkFunc) {
	for i := 1; i <= w.concurrentNum; i++ {
		go doWork(wf, w.jobsCh, w.resultCh)
		//fmt.Println("worker ", i, " initialized")
	}
}

func (w *worker) BuildJobs() {
	for i := 0; i < w.totalReqNum; i++ {
		w.jobsCh <- struct{}{}
		//fmt.Println("add job ", i+1)
	}
}

func (w *worker) PrintStatistic() {
	totalRespTime := int64(0)
	nfpRespTime := int64(0)
	nfpCount := int(float64(w.totalReqNum) * 0.95)
	for i := 0; i < w.totalReqNum; i++ {
		t := <-w.resultCh
		totalRespTime += t
		if nfpCount >= i {
			nfpRespTime += t
		}
	}

	fmt.Println("")
	fmt.Printf("avg response time:\t%.2Fs\n", float64(totalRespTime)/float64(w.totalReqNum)/float64(1000000000))
	fmt.Printf("95%% response time:\t%.2Fs\n", float64(nfpRespTime)/float64(nfpCount)/float64(1000000000))
	fmt.Println("")
}

func doWork(wf WorkFunc, jobs <-chan struct{}, respTimeCh chan<- int64) {
	for range jobs {
		respTime := wf.DoWork()
		respTimeCh <- respTime
		//fmt.Println("resp time: ", respTime)
	}
}
