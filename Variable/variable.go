package variable

import "fmt"

func Variable() {
	var name string

	name = "lemoncode21"
	fmt.Println(name)

	name = "youtube"
	fmt.Println(name)

	// write variable without data type
	var name2 = "my pc"
	fmt.Println(name2)

	name2 = "my old pc"
	fmt.Println(name2)

	// write variable without var
	name3 := "my youtube"
	fmt.Println(name3)

	name3 = "hello youtube"
	fmt.Println(name3)

	// Multiple Variable Declarations
	var (
		firstName = "lemoncode21"
		lastName  = "youtube"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Exercise
	fmt.Println("--Question 1--")
	var age int64
	age = 25
	fmt.Printf("Age: %d\n", age)

	fmt.Println("--Question 2--")
	var (
		name4   = "John"
		country = "USA"
	)
	fmt.Printf("Name: %s, Country: %s\n", name4, country)

	fmt.Println("--Question 3--")
	temperature := 25
	fmt.Printf("Temperature: %d\n", temperature)

}
