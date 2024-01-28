package datatype

import "fmt"

func BoolType() {
	//Boolean Data Type
	fmt.Println("True", true)
	fmt.Println("False", false)

	//Exercise
	fmt.Println("--Exercise--")
	var isLoggedIn bool
	fmt.Printf("Is the user currently logged in ? %t\n", isLoggedIn)
}
