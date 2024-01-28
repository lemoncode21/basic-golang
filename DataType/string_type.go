package datatype

import "fmt"

func StringType() {
	// string data type
	fmt.Println("Lemoncode21")
	fmt.Println("lemoncode21 youtube")

	fmt.Println(len("lemoncode21"))
	fmt.Println("lemoncode21"[0])
	fmt.Println("lemoncode21 youtube"[1])

	// Exercise
	fmt.Println("Exercise")
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Printf("Full Name:  %s\n", fullName)
}
