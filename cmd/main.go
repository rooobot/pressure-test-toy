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

	u := pkg.NewUrlWorkFunc(*targetURL)
	w := pkg.NewWorker(*targetURL, *concurrentNum, *totalReqNum)

	w.BuildWorker(u)
	w.BuildJobs()
	w.PrintStatistic()
}
