package ifelse

import "fmt"

func IfElse() {

	name := "John"

	if name == "John" {
		fmt.Println("Hello john")
	}

	// else

	if name == "John" {
		fmt.Println("Hello john")
	} else {
		fmt.Println("Hello Everyone")
	}

	// else if
	if name == "John" {
		fmt.Println("Hello john")
	} else if name == "David" {
		fmt.Println("Hello David")
	} else {
		fmt.Println("Hello Everyone")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("your name very long")
	} else {
		fmt.Println("name length is valid!")
	}

	//Exercise

	/*
		1. Declare a variable score and assign a value.
		Use an if-else statement to check if the score is greater than or equal to 60.
		Print "Pass" if true, otherwise print "Fail".
	*/

	fmt.Println("----Question 1 ----")
	score := 55
	if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	/*
		2. Declare two variables, isSunny and temperature.
		Use nested if-else statements to check if it's sunny and
		if the temperature is greater than 25 degrees Celsius. Print appropriate messages.
	*/

	fmt.Println("----Question 2 ----")
	isSunny := true
	temperature := 28

	if isSunny {
		if temperature > 25 {
			fmt.Println("It is a hot and sunny day!")
		} else {
			fmt.Println("It's sunny, but not very hot")
		}
	} else {
		fmt.Println("It's not sunny today")
	}
}
