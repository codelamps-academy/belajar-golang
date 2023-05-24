package main

import "fmt"

func main() {

	data1 := "Andre"
	data2 := "Andre"

	result := data1 == data2
	fmt.Println(result)
}

type MyClass struct {
	property1 string
	property2 int
}
