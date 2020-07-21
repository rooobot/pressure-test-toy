package main

import (
	"flag"
	"fmt"
	"net/url"
	"pressure-test-toy/pkg"
)

var (
	targetURL     = flag.String("url", "", "target URL for pressure test")
	concurrentNum = flag.Int("concurrentNum", 1, "concurrency number")
	totalReqNum   = flag.Int("totalReqNum", 1, "total request number")
)

func main() {
	flag.Parse()
	if *targetURL == "" {
		flag.Usage()
		return
	}

	_, err := url.ParseRequestURI(*targetURL)
	if err != nil {
		fmt.Println("invalid target url: ", *targetURL)
		return
	}

	u := &pkg.UrlWorkFunc{URL: *targetURL}

	w := &pkg.Worker{
		URL:           *targetURL,
		ConcurrentNum: *concurrentNum,
		TotalReqNum:   *totalReqNum,
		JobsCh:        make(chan struct{}, *totalReqNum),
		ResultCh:      make(chan int64, *totalReqNum),
	}

	w.BuildWorker(u)

	w.BuildJobs()

	w.PrintStatistic()
}
