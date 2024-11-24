package main

import "fmt"

type Coordinate struct {
	x *float64
	y *float64
}
type Shape struct {
	name *string
}

type Circle struct {
	*Shape
	radius *float64
}

type Rectangle struct {
	*Shape
	*Coordinate
}

func main() {

	circleName := "Circle"
	circleRadius := 3.14
	circle := Circle{Shape: &Shape{name: &circleName}, radius: &circleRadius}

	rName := "Rectangle"
	x, y := 12.1, 14.1
	rectangle := Rectangle{Shape: &Shape{name: &rName}, Coordinate: &Coordinate{x: &x, y: &y}}

	fmt.Println(circle, rectangle)
}
