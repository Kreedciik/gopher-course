package sum

import (
	"sync"
)

func sum(a int, ch chan int, wg *sync.WaitGroup) {
	start := (a-1)*100 + 1
	end := a * 100
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	ch <- sum
	wg.Done()
}

func Calculate() int {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	totalSum := 0

	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go sum(i, ch, &wg)
	}
	wg.Wait()
	close(ch)

	for v := range ch {
		totalSum += v
	}

	return totalSum

}
