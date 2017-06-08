package metric

import "time"
var t time.Time

func init() {
	t = time.Now()
}

func  Duration() int64 {
	return int64(time.Now().Sub(t))
}
