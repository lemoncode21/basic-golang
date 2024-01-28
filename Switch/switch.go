package Switch

import (
	"fmt"
	"time"
)

func Switchs() {
	name := "John"

	switch name {
	case "David":
		fmt.Println("Hello david")
	case "James":
		fmt.Println("Hello James")
	default:
		fmt.Println("what's your name!")
	}

	// swithc conditional case statement
	today := time.Now()
	fmt.Printf("what is today ? %v\n", today.Day())
	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house")

	case today.Day() <= 10:
		fmt.Println("Buy some snack")

	default:
		fmt.Printf("No information available for that day")
	}

	//Exercise

	/*
		1.Declare a variable grade and assign a letter grade (A, B, C, D, or F).
		Use a switch statement to print a message based on the grade
		(e.g., "Excellent", "Good", "Average", "Poor", "Fail").
	*/

	grade := "C"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good Job!")
	case "C":
		fmt.Println("Average.")
	case "D":
		fmt.Println("Poor.")
	case "F":
		fmt.Println("Fail.")
	default:
		fmt.Println("Invalid grade.")
	}

}
