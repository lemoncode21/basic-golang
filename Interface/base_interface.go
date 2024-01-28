package Interface

import (
	"fmt"
	"math"
)

type Employee interface {
	PrintName(name string)
}

type Attributes struct {
	id       int
	username string
}

func (a Attributes) PrintName(name string) {
	fmt.Println("Employee Id:\t", a.id)
	fmt.Println("Employee username:\t", a.username)
	fmt.Println("Employee Name:\t", name)
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Ups() interface{} {
	// return 1
	// return "String"
	return true
}

func BaseInterface() {
	var e1 Employee
	e1 = Attributes{id: 23, username: "user23"}
	e1.PrintName("David")

	// empty interface type
	empty := Ups()
	fmt.Println(empty)

	//Exercise

	/*
		1.Define an interface called Shape with a method Area() that returns a float64.
		Implement this interface with two types: Circle and Rectangle.
		Create instances of each type, call the Area() method, and print the results.
	*/

	circle := Circle{Radius: 3}
	rectangle := Rectangle{Width: 4, Height: 5}

	fmt.Printf("Circle Area : %f\n", circle.Area())
	fmt.Printf("Rectangle Area: %f\n", rectangle.Area())
}
