package main

import "fmt"

// func divide(a, b float64) float64 {
// 	return a / b
// }

// func main() {
// 	fmt.Println("Starting the program")
// 	solution := divide(10, 3)
// 	fmt.Println(solution)
// }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func main() {

	// at the  place of the error we can use the  underscore
	// underscore handle the errors efficently
	ans, _ := divide(10, 0)
	fmt.Println(ans)

}
