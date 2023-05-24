package main

import "fmt"

func myNewFunction(firstName string, lastName string) string {
	if firstName == "" {
		return "Hello Bro"
	} else {
		return "Hello " + firstName + lastName
	}
}

func main() {
	result := myNewFunction("Andre Rizaldi", "Brillianto")
	result2 := myNewFunction("", "Budiono")
	fmt.Println(result)
	fmt.Println(result2)
}
