package main

import (
	"fmt"
	"time"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Printf("recovering from the panic %w", r)
		fmt.Println()
	}
}

func doSomething() {
	defer recoverFromPanic()
	fmt.Println("Do something")
	panic("Do smth panic")
}
func main() {
	now := time.Now()
	fmt.Println(now.Format("02.01.2006 03:04"))
	// DD.MM.YYYY
}
