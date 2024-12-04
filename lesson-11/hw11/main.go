package main

import (
	"fmt"
	n "hw11/numbers"
	b "hw11/package1"
	p2 "hw11/package2"
	c "hw11/sum"
	"sync"
)

// 3.1

func main() {
	// 3.1
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go n.PrintNumbers(&wg1)
	wg1.Wait()

	// 3.2
	totalSum := c.Calculate()
	fmt.Println("The sum of numbers between [1, 1000]: ", totalSum)

	// 3
	books := b.GetBooks()
	var wg sync.WaitGroup
	ch := make(chan b.Book, 5)

	wg.Add(5)
	ch <- books[0]
	ch <- books[1]
	ch <- books[2]
	ch <- books[3]
	ch <- books[4]

	p2.WriteToJSON(ch, &wg)
	p2.WriteToJSON(ch, &wg)
	p2.WriteToJSON(ch, &wg)
	p2.WriteToJSON(ch, &wg)
	p2.WriteToJSON(ch, &wg)
	wg.Wait()

	channelForRead := make(chan b.Book, 5)
	p2.ReadFromJSON(channelForRead)
	close(channelForRead)
	for book := range channelForRead {
		fmt.Println(book)
	}
}
