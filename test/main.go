package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			ch1 <- "Half a second passed"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "2 seconds passed"
		}
	}()

	for {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
		// select {
		// case msg := <-ch1:
		// 	fmt.Println(msg)
		// case msg := <-ch2:
		// 	fmt.Println(msg)
		// }
	}
}
