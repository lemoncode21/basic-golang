package constant

import "fmt"

func Constants() {
	const firstName string = "lemoncode21"
	const lastName = "youtube"

	// error
	// firstName = "change name"
	// lastName = "change name"

	const (
		firstName2 string = "my pc"
		lastName2         = "my old pc"
	)

	// firstName2 = "cannot change constant"
	// lastName2 = "cannot change constant"

	// Exercise
	fmt.Println("---Question 1---")
	const pi = 3.14159
	fmt.Println(pi)

	fmt.Println("---Question 2---")
	const (
		daysInWeek = 7
		hoursInDay = 24
	)
	fmt.Println(daysInWeek)
	fmt.Println(hoursInDay)
}
