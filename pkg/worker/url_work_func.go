package worker

import (
	"fmt"
	"net/http"
	"time"
)

type urlWorkFunc struct {
	url string
}

func NewUrlWorkFunc(url string) WorkFunc {
	return &urlWorkFunc{url: url}
}

func (u *urlWorkFunc) DoWork() int64 {
	start := time.Now().UnixNano()
	_, err := http.Get(u.url)
	if err != nil {
		fmt.Println(err)
	}
	//time.Sleep(1 * time.Second)
	end := time.Now().UnixNano()

	return end - start
}
