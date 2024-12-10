package assignment6

import (
	"time"
)

func TimeDistance(ch chan time.Time, second time.Time) {
	t := <-ch
	diff := second.Sub(t)
	elapsed := time.Now().Add(diff)
	ch <- elapsed
}
