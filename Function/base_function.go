package Function

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func add(x int, y int) {
	total := 0
	total = x + y

	fmt.Println(total)
}

func addReturn(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return
}

func getFullName() (string, string) {
	return "John", "David"
}

func getComplateName() (firstName, lastName string) {
	firstName = "James"
	lastName = "David"

	return firstName, lastName
}

type Filters func(string) string

func sayHelloWithFilter(name string, filter Filters) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "dog" {
		return "..."
	} else {
		return name
	}
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func rectangleInfo(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func counterWithState() func(increment int) int {
	count := 0
	return func(increment int) int {
		count += increment
		return count
	}

}

func processNumbers(numbers []int, callback func(int) int) {
	for _, num := range numbers {
		result := callback(num)
		fmt.Println(result)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func MainFunction() {
	// function
	sayHello()

	// function with parameters
	add(20, 30)

	// function with return value
	sum := addReturn(10, 30)
	fmt.Println(sum)

	// function return multiple value
	var a, p int
	a, p = rectangle(10, 30)
	fmt.Println("Area: ", a)
	fmt.Println("Parameter: ", p)

	//funtion ignore return value
	firstName, _ := getFullName()

	fmt.Println(firstName)

	// function named return value
	firstName, lastName := getComplateName()
	fmt.Println(firstName, lastName)

	// function as parameter
	sayHelloWithFilter("John", spamFilter)
	sayHelloWithFilter("dog", spamFilter)

	// Anonymous Function
	area := func(l int, b int) int {
		return l * b
	}

	fmt.Println(area(20, 30))

	// recrusive function
	fmt.Println("---Recursive---")
	fmt.Println(factorialRecursive(5))

	// closure
	fmt.Println("---Closure---")
	l := 20
	b := 12

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()

	// Exercise

	/*
		1.Write a function called rectangleInfo that takes the length and width of a rectangle as parameters
		and returns the area and perimeter with named return values.
	*/
	fmt.Println("--Question 1--")
	var areas, perimeter float64
	areas, perimeter = rectangleInfo(4, 5)
	fmt.Println("Area :", areas)
	fmt.Println("Perimeter : ", perimeter)

	/*
		2.Write a function called counterWithState that returns a closure.
		The closure should maintain a counter variable and also accept an increment value as a parameter.
	*/
	fmt.Println("--Question 2--")
	counterFunc := counterWithState()
	fmt.Println(counterFunc(2), counterFunc(3), counterFunc(5))

	/*
		3.Write a function called processNumbers that takes a slice of integers and a callback function.
		Apply the callback function to each element of the slice.
	*/
	fmt.Println("--Question 3--")
	intSlice := []int{1, 2, 3, 4, 5}
	processNumbers(intSlice, func(x int) int {
		return x * x
	})

	// Exericise Anonymous Function

	/*
		1. Declare a variable multiply and assign an anonymous function to it.
		The function should take two integers as parameters and return their product.
	*/

	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(2, 4)
	fmt.Println("Product: ", product)

	// Exercise

	/*
		1. Write a recursive function called fibonacci that takes a non-negative integer as a parameter
		and returns the nth Fibonacci number.
	*/
	result := fibonacci(7)
	fmt.Println(result)
}
