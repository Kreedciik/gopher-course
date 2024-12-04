package numbers

import (
	"fmt"
	"sync"
)

func PrintNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
