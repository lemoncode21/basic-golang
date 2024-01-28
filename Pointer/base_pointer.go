package Pointer

import "fmt"

type Employee struct {
	Id            int
	Name, Address string
}

func ChangeAddressToNevada(employee *Employee) {
	employee.Address = "Nevada"
}

func (e *Employee) Married() {
	e.Name = "Mr. " + e.Name
}

type Car struct {
	Brand string
	Price float64
}

func PointerBase() {
	employee1 := Employee{1, "John", "Chicago"}
	employee2 := &employee1

	employee2.Address = "New York"

	*employee2 = Employee{2, "Dadvid", "Tokyo"}

	fmt.Println(employee1)
	fmt.Println(employee2)

	// pointer new
	fmt.Println("---pointer new ---")
	e1 := new(Employee)
	e2 := e1

	e2.Address = "New York"

	fmt.Println(e1)
	fmt.Println(e2)

	// pointer function
	fmt.Println("---pointer function---")
	em1 := Employee{3, "Norman", "San Fransisco"}
	ChangeAddressToNevada(&em1)

	fmt.Println(em1)

	// pointer in method
	fmt.Println("---pointer in method---")
	ep1 := Employee{4, "David", "London"}
	ep1.Married()
	fmt.Println(ep1)

	//Exercise

	/*
		1. Define a struct called Car with fields for Brand (string) and Price (float64).
		Declare an array of Car structs. Declare an array of pointers to Car structs
		and assign each element a pointer to a Car struct from the original array.
		Print the values of the cars through the pointers.
	*/

	cars := []Car{
		{Brand: "Toyota", Price: 25000},
		{Brand: "Honda", Price: 27000},
		{Brand: "Ford", Price: 30000},
	}

	var carPointers [3]*Car
	for i := range cars {
		carPointers[i] = &cars[i]
	}

	for _, carPointer := range carPointers {
		fmt.Printf("Brand: %s, Price: %.2f\n", carPointer.Brand, carPointer.Price)
	}
}
