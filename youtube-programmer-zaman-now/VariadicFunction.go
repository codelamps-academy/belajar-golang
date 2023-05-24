package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	// using slice
	numbers := []int{2, 2, 2, 2, 2}
	total = sumAll(numbers...)
	fmt.Println(total)
}
