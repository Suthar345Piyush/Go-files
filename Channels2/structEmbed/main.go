package main

import "fmt"

//struct embedding

type Animal struct {
	Name string
	Age  int
}

type Pet struct {
	Animal // this  is  struct embedding
	Owner  string
}

func main() {

	pet := Pet{
		Animal: Animal{Name: "Shinti", Age: 2},
		Owner:  "Piyush",
	}

	fmt.Println(pet.Name)
	fmt.Println(pet.Age)
	fmt.Println(pet.Owner)

}
