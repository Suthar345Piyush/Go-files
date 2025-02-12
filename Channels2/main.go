<<<<<<< HEAD
package main

import (
	"fmt"
	"mylearning/myUtil"
)

func main() {
	fmt.Println("Hello i am learning Golang")
	fmt.Println("Using golang to excel in tech")
	fmt.Println("I am frontend dev")
	fmt.Println("I am a second year ug student")

	myUtil.PrintMessage("hello bois")
=======
//Channels  deep dive , with more examples
//Big code example

package main

import "fmt"

func calcSquare(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}

	squareop <- sum
}

func calcCube(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum = sum + digit*digit*digit
		number /= 10
	}

	cubeop <- sum
}

func main() {
	number := 50

	//square channel

	sqrch := make(chan int)
	//cube channel

// 	cbrch := make(chan int)

	// goroutine of square sum

	go calcSquare(number, sqrch)

	//goroutine of cube sum

	go calcCube(number, cbrch)

	//receiving the data
	squares, cube := <-sqrch, <-cbrch

	fmt.Println("Final summation is: ", squares+cube)

}

//DEADLOCK IN GOLANG

//If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.

package main

func main() {
	ch := make(chan int)
	ch <- 5
}

// unidirectional channels
// these channels only sends  or only receive the  data

package main

import "fmt"

//this  chan<- int denotes the send only channel
func newFunc(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)
	go newFunc(sendch)
	fmt.Println(<-sendch)
}

//bidirectional channel to unidirectional channel is  can be done , but the vice-verca is not possible

//for example

package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
>>>>>>> a2a782e (Pushing go files)
}
