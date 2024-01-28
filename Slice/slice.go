package slice

import "fmt"

func Slices() {
	var strSlice = []string{"USA", "UK", "JAPAN"}

	fmt.Printf("strSlice: %s", strSlice)
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

	// Declare Slice using make
	var strSlice2 = make([]string, 10, 20)

	fmt.Printf("strSlice2 \tLen: %v \tCap: %v\n", len(strSlice2), cap(strSlice2))

	// Access Items
	var intSlice = []int{10, 20, 30, 40}
	fmt.Println(intSlice[0])
	fmt.Println(intSlice[1])
	fmt.Println(intSlice[0:4])

	// add slice
	strSlice = append(strSlice, "germany")
	fmt.Println(strSlice)

	// copy slice
	strSlice3 := make([]string, 5, 10)
	copy(strSlice3, strSlice)

	fmt.Println("After Copying: ", strSlice3)

	//Exercise

	/*
		1. Declare a slice named numbers with integer values 1 to 5. Print the slice.
	*/
	fmt.Println("----Question 1 ----")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers: %v\n", numbers)

	/*
		2. Declare a slice named colors with string values representing colors.
		Append a new color to the slice and print the updated slice.
	*/
	fmt.Println("----Question 2 ----")

	colors := []string{"Red", "Green", "Blue", "Yellow"}
	colors = append(colors, "Purple")
	fmt.Printf("Updated Colors: %v\n", colors)
}
