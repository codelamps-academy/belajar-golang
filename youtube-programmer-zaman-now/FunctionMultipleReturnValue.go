package main

import "fmt"

func myMultipleReturnValueFunction(firstName string, lastName string) (string, string) {
	name := firstName + " " + lastName
	return "Hello", name
}

func myMethod() (string, string) {
	return "Andre ", "Rizaldi"
}

func main() {
	println(myMultipleReturnValueFunction("Andre", "Rizaldi"))

	firstName, _ := myMethod()
	fmt.Println(firstName)

	_, lastName := myMethod()
	fmt.Println(lastName)
}
