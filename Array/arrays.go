package array

import "fmt"

func Array() {
	var theArray [3]string
	theArray[0] = "USA"
	theArray[1] = "Japan"
	theArray[2] = "UK"

	fmt.Println(theArray[0])
	fmt.Println(theArray[1])
	fmt.Println(theArray[2])

	// Initializing an array with an Array Literal
	x := [5]int{1, 2, 3, 4, 5}

	fmt.Println(x)

	// Intializing an array with ellipses...
	y := [...]int{1, 2, 3}

	fmt.Println(y)

	//Exercise

	/*
		1. Declare an array named numbers with integer values 1 to 5. Print the array
	*/

	fmt.Println("----Question 1 ----")
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Numbners: %v\n", numbers)

	/*
		2. Declare an array named fruits with three string values representing fruits.
		Print the second element of the array.
	*/

	fmt.Println("----Question 2 ----")
	fruits := [3]string{"Apple", "Banana", "Orange"}
	fmt.Printf("Second Fruits: %s\n", fruits[1])

	/*
		3. Declare an array named ages with integer values.
		Modify the third element of the array to a new value and print the updated array.
	*/
	fmt.Println("----Question 3 ----")
	ages := [4]int{25, 30, 35, 40}
	ages[2] = 36
	fmt.Printf("Updated Ages: %v\n", ages)
}
