package main

import "fmt"

func SayHello() {
	fmt.Println("Hello  Andre")
}

func MyFunction() string {
	name := "Andre Rizaldi Brillianto"
	return name
}

func main() {
	SayHello()
	SayHello()
	SayHello()

	println(MyFunction())
	println(MyFunction())
	println(MyFunction())
}
