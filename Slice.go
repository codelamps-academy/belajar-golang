package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	// mengakses elemen slice menggunakan indeks
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])

	// MENGUBAH NILAI ELEMEN SLICE
	numbers[1] = 10
	fmt.Println(numbers[1])

	numbers = append(numbers, 6)
	fmt.Println(numbers)
}
