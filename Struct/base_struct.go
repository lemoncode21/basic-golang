package Struct

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) printDetails() {
	fmt.Printf("Name: %s, Age: %d\n", c.Name, c.Age)
}

type Person struct {
	Name, Address string
	Age           int
}

func (p Person) PrintDetails() {
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.Name, p.Age, p.Address)
}

type Address struct {
	City    string
	Country string
}

type PersonWithAddress struct {
	Person
	Address
}

func BaseStruct() {
	var customer Customer
	customer.Name = "John"
	customer.Address = "Jl Adline no 3"
	customer.Age = 25

	fmt.Println(customer)

	newCustomer := Customer{
		Name:    "David",
		Address: "USA",
		Age:     34,
	}

	fmt.Println(newCustomer)

	secondCustomer := new(Customer) //is a pointer instance
	secondCustomer.Name = "Taylor"
	secondCustomer.Address = "USA"
	secondCustomer.Age = 28

	fmt.Println(secondCustomer)

	var secondCustomer2 = new(Customer)
	secondCustomer2.Name = "James"
	secondCustomer2.Address = "USA"
	secondCustomer2.Age = 25

	fmt.Println(secondCustomer2)
	secondCustomer2.printDetails()

	// Exercise

	/*
		1.Define a struct called Address with fields for City and Country.
		Extend the Person struct to include an Address field.
		Create a person with an address and print their details.
	*/

	personWithAddress := PersonWithAddress{
		Person: Person{Name: "David", Age: 25, Address: "riverside"},
		Address: Address{
			City:    "New york",
			Country: "USA",
		},
	}
	personWithAddress.PrintDetails()
	fmt.Println(personWithAddress)
}
