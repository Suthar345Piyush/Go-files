package main

import "fmt"

func main() {
	age := 19
	year := 2005
	name := "piyush"

	fmt.Println("age:", age, "year:", year, "name:", name)

	//the printf function is used in go for formating the data  in a  better way

	// fmt.Printf("The age is:%d\n", age)
	// fmt.Printf("The year is : %d\n", year)
	// fmt.Printf("The name is : %s\n", name)

	fmt.Printf("Age is  just: %d\n", age)
	fmt.Printf("Name: %s , Age : %d , year : %d", name, age, year)

}
