package main

import "fmt"

func namedReturnValue() (firstName, middleName, lastName string) {
	firstName = "Andre"
	middleName = "Rizaldi"
	lastName = "Brillianto"

	return
}

func main() {
	firstName, _, lastName := namedReturnValue()
	fmt.Println(firstName)
	fmt.Println(lastName)
}