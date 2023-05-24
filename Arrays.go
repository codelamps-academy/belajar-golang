package main

import "fmt"

func main() {

	// ARRAY MANUAL
	var data [4]string
	data[0] = "Andre"
	data[1] = "Rizaldi"
	data[2] = "Brillianto"
	data[3] = "Junior"

	fmt.Println(data[0])
	fmt.Println(data[1])
	fmt.Println(data[2])
	fmt.Println(data[3])

	var values = [4]int{
		10,
		20,
		30,
		40,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
}
