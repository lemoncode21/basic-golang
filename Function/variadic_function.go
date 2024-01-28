package Function

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func calculateAverage(numbers ...float64) (average float64) {
	sum := 0.0
	count := 0.0
	for _, num := range numbers {
		sum += num
		count++
	}
	average = sum / float64(count)
	return
}

func VariadicFunction() {
	//variadic function
	total := sumAll(10, 11, 12, 13)
	fmt.Println(total)

	// input slice parameter
	numbers := []int{20, 21, 22, 23}
	total = sumAll(numbers...)
	fmt.Println(total)

	// Exericise

	/*
		1. Write a variadic function called calculateAverage that
		takes an arbitrary number of floats as parameters
		and returns the average with a named return value.
	*/

	averageResult := calculateAverage(3.0, 4.5, 6.0, 7.8)
	fmt.Println("Average: ", averageResult)

}
