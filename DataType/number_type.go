package datatype

import (
	"fmt"
	"math"
)

func NumberType() {

	// Number Data Type
	fmt.Println("One = ", 1)
	fmt.Println("two = ", 2)
	fmt.Println("three point five = ", 3.5)

	//Exercisse Number Data Type
	// Question 1
	fmt.Printf("Question 1")
	integerValue := 42
	floatValue := float64(integerValue)
	fmt.Printf("Conversion of %d to float64: %f\n", integerValue, floatValue)
	//Question 2
	fmt.Printf("Question 2")
	pi := 3.1415
	roundedPi := math.Round(pi)
	fmt.Printf("Rounded Pi: %.0f\n", roundedPi)
}
