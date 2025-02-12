package main

import "fmt"

func main() {
	fmt.Println("Datatypes in Golang")

	var name string = "piyush"
	fmt.Println(name)

	var year = 2024
	fmt.Println(year)

	var month = "december"
	fmt.Println(month)

	const month2 string = "december"
	fmt.Println(month2)

	var decide bool = true
	fmt.Println(decide)

	const december bool = true
	fmt.Println(december)

	// name := "piyush"
	// fmt.Println(name)

	var Public = "data can be shared"
	var private = "data is private"

	var complex64Num complex64 = 1 + 2i
	var complex128Num complex128 = 1 + 3.14i

	fmt.Println(complex64Num)
	fmt.Println(complex128Num)

	fmt.Println(Public)
	fmt.Println(private)

}
