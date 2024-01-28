package Map

import "fmt"

func Maps() {

	var employee = map[string]int{"mark": 10, "sandy": 10}

	fmt.Println(employee)

	// make declaration using make function
	var newEmployee = make(map[string]int)

	newEmployee["Mark"] = 12
	newEmployee["John"] = 10

	fmt.Println(newEmployee)

	// Accesing item
	fmt.Println(newEmployee["Mark"])

	// Adding Item
	newEmployee["David"] = 8
	fmt.Println(newEmployee)

	// Update Item
	newEmployee["John"] = 100
	fmt.Println(newEmployee)

	// Delete Item
	delete(newEmployee, "Mark")
	fmt.Println(newEmployee)

	//Exercise

	/*
		1. Declare a map named countries where keys are country names (strings)
		and values are their populations (integers). Add new country,  Modify the populations some country
		and Delete an entry for a specific country and print the updated map.
	*/

	countries := map[string]int{
		"USA":    331,
		"China":  1440,
		"India":  1380,
		"Brazil": 213,
	}

	// Add
	countries["Japan"] = 123
	fmt.Printf("Add Country : %v\n", countries)

	// Update
	countries["China"] = 1550
	fmt.Printf("Update Country : %v\n", countries)

	// Delete
	delete(countries, "Brazil")
	fmt.Printf("Delete Country : %v\n", countries)

}
