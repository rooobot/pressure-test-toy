package pkg

import "fmt"

type Worker struct {
	URL           string
	TotalReqNum   int
	ConcurrentNum int
	JobsCh        chan struct{}
	ResultCh      chan int64
}

type WorkFunc interface {
	DoWork() int64
}

func (w *Worker) BuildWorker(wf WorkFunc) {
	for i := 1; i <= w.ConcurrentNum; i++ {
		go doWork(wf, w.JobsCh, w.ResultCh)
		//fmt.Println("worker ", i, " initialized")
	}
}

func (w *Worker) BuildJobs() {
	for i := 0; i < w.TotalReqNum; i++ {
		w.JobsCh <- struct{}{}
		//fmt.Println("add job ", i+1)
	}
}

func (w *Worker) PrintStatistic() {
	totalRespTime := int64(0)
	nfpRespTime := int64(0)
	nfpCount := int(float64(w.TotalReqNum) * 0.95)
	for i := 0; i < w.TotalReqNum; i++ {
		t := <-w.ResultCh
		totalRespTime += t
		if nfpCount >= i {
			nfpRespTime += t
		}
	}

	fmt.Println("")
	fmt.Printf("avg response time:\t%.2Fs\n", float64(totalRespTime)/float64(w.TotalReqNum)/float64(1000000000))
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
