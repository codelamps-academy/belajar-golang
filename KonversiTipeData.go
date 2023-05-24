package main

import "fmt"

func main() {

	var data1 int8 = 10
	var data2 int16 = int16(data1)
	var data3 int32 = int32(data2)

	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)

}
