package operation

import "fmt"

func Operations() {
	// aritmatic Operation
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	// assignment operation
	var x, y = 15, 25
	x = y

	fmt.Println("=", x)

	x = 15
	x += y

	fmt.Println("+=", x)

	x = 50
	x -= y
	fmt.Println("-=", x)

	// comparison Operation

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x <= y)
	fmt.Println(x >= y)
	fmt.Println(x > y)
	fmt.Println(x < y)

	//Exercise

	/*
		1. Declare two variables, a and b, with values of your choice.
		Perform and print the results of addition, subtraction, multiplication, and division.
	*/

	fmt.Println("---Question 1 ---")
	a, b = 10, 5
	fmt.Printf("Addition: %d\n", a+b)
	fmt.Printf("Subtraction: %d\n", a-b)
	fmt.Printf("Multiplication: %d\n", a*b)
	fmt.Printf("Division: %d\n", a/b)

	/*
		2. Declare a variable total with an initial value.
		Use assignment operations to add, subtract, multiply, and divide the variable by another value.
		Print the variable after each operation.
	*/

	fmt.Println("---Question 2---")
	total := 50

	total += 10
	fmt.Printf("After addition: %d\n", total)

	total -= 5
	fmt.Printf("After subtraction: %d\n", total)

	total *= 2
	fmt.Printf("After multiplication: %d\n", total)

	total /= 3
	fmt.Printf("After division: %d\n", total)

	/*
		3. Declare two boolean variables, isSunny and isWarm, with values of your choice.
		Use logical operations (AND, OR, NOT) to determine if it's a good day for outdoor activities.
		Print the results.
	*/

	isSunny, isWarm := true, true
	fmt.Printf("Is it good for outdoor activities ? %t\n", isSunny && isWarm)
	fmt.Printf("Is it good for indoor activities? %t\n", isSunny || isWarm)
	fmt.Printf("Is it not good for outdoor activities? %t\n", !isSunny)

}
