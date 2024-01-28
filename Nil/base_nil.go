package Nil

import "fmt"

func printIfNotNil(ptr *int) {
	if ptr != nil {
		fmt.Println("Value: ", *ptr)
	} else {
		fmt.Println("Pointer is nil")
	}
}

func BaseNil() {
	printIfNotNil(nil)
	value := 12
	printIfNotNil(&value)
}
