//generics  in golang

package main

import "fmt"

//heres any means  it can be of any time
//we cn write  interface{} at the  place  of the  any

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printSliceString(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func main() {

	//nums := []int{1 ,2, 3, 4, 5}
	names := []string{"a", "b", "c", "d"}
	names2 := []bool{true, false, true, false}

	// names := []int{1, 2, 3, 79090, 8}
	printSlice(names)
	printSlice(names2)

	// printSliceString(names)
}

//notes :
// we have  to write multiple  functions with some  changes  in the  parameters , so avoid  this  problem we  can use  the  generics  in golang
