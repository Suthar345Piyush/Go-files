package main

import "fmt"

func firstFunction() {
	fmt.Println("this  is  our  first function")
}

func addNumbers(a, b int) int {
	return a + b
}

// we can write this in that way also

func multiplyNum(a, b int) (result int) {
	result = a * b
	return

}

func main() {
	fmt.Println("i am learning golang")
	firstFunction()

	solution := addNumbers(10, 5)
	fmt.Println("The sum of a and b is: ", solution)

	multiply := multiplyNum(10, 5)
	fmt.Println("the multiply of a and b is: ", multiply)
}
