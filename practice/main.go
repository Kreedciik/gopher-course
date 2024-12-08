package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait(c chan int, wg *sync.WaitGroup) {
	randomValue := rand.Intn(5)
	time.Sleep(time.Duration(randomValue) * time.Second)
	wg.Done()
	c <- randomValue
}
func main() {
	stopChan := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		stopChan <- true
	}()
	timer := time.NewTimer(time.Second)

LOOP:
	for {
		select {
		case <-timer.C:
			fmt.Println("Tick")
			timer.Reset(time.Second)
		case <-stopChan:
			fmt.Println("BOOM")
			break LOOP
		}
	}

}
