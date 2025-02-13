package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id       int
	Title    string
	Duration time.Duration
}

func doTask(task Task) {
	fmt.Printf("Started task #%d \n", task.Id)
	time.Sleep(task.Duration)
	fmt.Printf("Finished task #%d \n", task.Id)
}

func main() {
	var wg sync.WaitGroup
	tasks := []Task{
		{1, "Cook soup", 3 * time.Second},
		{2, "Cook main meal", 5 * time.Second},
		{3, "Cook desert", 2 * time.Second},
	}

	for _, task := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doTask(task)
		}()
	}
	wg.Wait()
}
