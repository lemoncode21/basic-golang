package forloop

import "fmt"

func Loop() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Loop", counter)
		counter++
	}

	// loop with statement

	for ncounter := 1; ncounter <= 10; ncounter++ {
		fmt.Println("new loop", ncounter)
	}

	// loop with range
	name := []string{"James", "Mahmud", "Taylor"}

	for index, element := range name {
		fmt.Println("Index: ", index, "Element: ", element)
	}

	// loop only key
	for key := range name {
		fmt.Println(key)
	}

	// loop only value
	for _, value := range name {
		fmt.Println(value)
	}

	//Exercise

	/*
		1. Use a for loop to print even numbers from 2 to 10.
		Use an if statement to determine if a number is even.
	*/

	fmt.Println("---question 1 ---")
	for i := 2; i <= 10; i += 2 {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	/*
		2. Declare an array of strings fruits with a few fruit names.
		Use a for loop with range to print each fruit.
	*/
	fmt.Println("---question 2 ---")
	fruits := []string{"Apple", "Banana", "Orange", "Grapes"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
