package main

import (
	"fmt"
	"os"
)

// Homework-4
func main() {
	var login, password string
	var correctLogin, correctPassword string = "login", "password"
	var attempts uint = 3

	for {
		fmt.Print("Login: ")
		fmt.Scan(&login)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		if login != correctLogin && password != correctPassword {
			attempts--
			if attempts == 0 {
				fmt.Println("You entered 3 times incorrectly")
				os.Exit(1)
			} else {
				fmt.Println("Login or password is incorrect!")
				fmt.Printf("You have %v attempts left", attempts)
				fmt.Println()
			}

		} else {
			fmt.Println("Access granted!")
			break
		}
	}
}
