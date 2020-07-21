package pkg

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_DoWork(t *testing.T) {
	u := &UrlWorkFunc{URL: "https://www.baidu.com"}
	latency := u.DoWork()
	assert.Equal(t, latency >= 0, true)
}
