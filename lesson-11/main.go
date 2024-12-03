package main

import (
	"fmt"
	"time"
)

// Go routines
func main() {
	// unbuffered channel
	ch := make(chan int)
	go sqr(ch)
	time.Sleep(3 * time.Second)
	ch <- 12
	fmt.Println(<-ch)
}

func sqr(ch chan int) {
	time.Sleep(2 * time.Second)
	a := <-ch
	ch <- a * a
}
