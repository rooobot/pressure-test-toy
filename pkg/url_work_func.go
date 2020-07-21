package pkg

import (
	"fmt"
	"net/http"
	"time"
)

type UrlWorkFunc struct {
	URL string
}

func (u *UrlWorkFunc) DoWork() int64 {
	start := time.Now().UnixNano()
	_, err := http.Get(u.URL)
	if err != nil {
		fmt.Println(err)
	}
	//time.Sleep(1 * time.Second)
	end := time.Now().UnixNano()

	return end - start
}
