//Channels  is  like pipe thorugh which the goroutines can communicate and can send data to each other
//chan T , is the  channel of type  T
//zero value is " nil "  , if a  channel is  nil , then it is useless and  , if  channels is using " make " keyword , works  same  like  maps  and  slices

// package main

// import "fmt"

// func main() {
// 	var a chan int

// 	if a == nil {

// 		fmt.Println("Our Channel is nil , making it useful")
// 		a := make(chan int)
// 		fmt.Printf("Type of a  %T", a)
// 	}
// }

//thats  how the channels  send & receive the data

// data := <- a , reading from a
// a  <- data  , writing into a

//Sends and receives are blocking bydefault

//In Go, by default, sends and receives on channels are blocking operations. This means that when a goroutine attempts to send a value to a channel, it will wait until another goroutine is ready to receive that value, and vice versa.

//example  of  goroutines how they communicate using channels

// package main

// import "fmt"

// func hello(done chan bool) {
// 	fmt.Println("hello world goroutine")
// 	done <- true
// }

// func main() {
// 	done := make(chan bool)
// 	go hello(done)
// 	<-done // receiving data from done channel
// 	fmt.Println("main function")
// }

//Example  with sleep method , for clear understanding

package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello goroutine with sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello from goutine awake")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main program to call the goroutine")
	go hello(done)
	<-done
	fmt.Println("main received data")
}
