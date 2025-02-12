//Generics  in golang

//Generics  also known as  parametric polymorphism
// it  enables us  to write code that operates on multiple  types without specifying  types
//generics means  jinka koi specific type nhi hota hai , koi specific type of  input/ouptut parameters nhi hota hai

// generics  in golang  has :
//1 . Type  parameters -> it  like  placeholder , in the function initialization
//2. Type  constraints -> restrict the  types, that can be  used as  type  arguments

// simple code :

// package main

// import "fmt"

// // T is a  type  parameter
// // any is  constraint , it  is  built in , any type can be  used
// func PrintSlice[T any](s []T) {
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

// func main() {
// 	intSlice := []int{1, 2, 3}
// 	stringSlice := []string{"hello", "world"}

// 	PrintSlice[string](stringSlice)
// 	PrintSlice[int](intSlice)
// }

// PROG : 2 finding  min in the  slice

// package main

// import (
// 	"fmt"

// 	"golang.org/x/exp/constraints"
// )

// func Min[T constraints.Ordered](s []T) T {

// 	//panic  stops  the  normal excution of the current  goroutine
// 	if len(s) == 0 {
// 		panic("empty slice")
// 	}

// 	min := s[0]
// 	for _, v := range s {
// 		if v < min {
// 			min = v
// 		}
// 	}
// 	return min
// }

// func main() {
// 	intSlice := []int{3, 1, 4, 2}
// 	floatSlice := []float64{3.14, 1.61, 4.67, 2.71}

// 	fmt.Println(Min[int](intSlice))
// 	fmt.Println(Min[float64](floatSlice))
// }

//PROG-3 Impelamenting a  data structure(stack)

package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return value, true
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func main() {

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Size())

	stringStack := Stack[string]{}
	stringStack.Push("piyush")
	stringStack.Push("suthar")

	fmt.Println(stringStack.Pop())
	fmt.Println(stringStack.Size())

}

// Generics in Go provide a powerful way to write reusable and type-safe code.
