package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	// 1 rune and byte
	// ? rune data type represents the unicode code point of character
	var r rune = '♄'
	fmt.Printf("%#U \n", r)
	fmt.Printf("%+q \n", r)
	// ? All string in Go consists of sequence of bytes
	str := "Golang"
	strToBytes := []byte(str)
	bytesToStr := string(strToBytes)
	fmt.Printf("%s => %v\n", str, strToBytes)
	fmt.Printf("%v => %s", strToBytes, bytesToStr)
	fmt.Println("\n")

	// 2.1 Displaying the full name
	fName := "Islamgaliev"
	lName := "Ilfat"
	fmt.Println(fName + " " + lName + "\n")

	// 2.2 Displaying the age
	birthYear := 1998
	age := time.Now().Year() - birthYear
	fmt.Println("Your age is", age)
	fmt.Println()
	// 2.3 Displaying the square of the numbers [1,10]
	for i := 1; i <= 10; i++ {
		fmt.Printf("The square of %v is %v \n", i, int(math.Pow(float64(i), 2)))
	}

	// 2.4 Math operations on float and int
	var a int32 = 20
	var b float64 = 5.0
	var c int32 = 6
	fmt.Println()

	// ? + operation
	sum := float64(a) + b
	fmt.Printf("%v + %.2f = %.2f\n", a, b, sum)

	// ? - operation
	subtract := float64(a) - b
	fmt.Printf("%v - %.2f = %.2f\n", a, b, subtract)

	// ? * operation
	product := float64(a) * b
	fmt.Printf("%v * %.2f = %.2f\n", a, b, product)

	// ? / operation
	intDivision := a / c
	fmt.Printf("%v / %v = %v \n", a, c, intDivision)

	// ? % operation
	residualDivision := a % c
	fmt.Printf("%v %% %v = %v\n", a, c, residualDivision)
}
