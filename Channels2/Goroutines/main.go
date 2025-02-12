//Goroutines  in Golang
//goroutines  are  light  weight  threads which is used for  multithreading & concurrency.
//Golang  is  concurrent language , it  does not  follow parallelism.

//to use them just write " go " , in the front  of  function invocation , then it  will work like a light weight thread

//this basically depands on the  core the  machine

//parallel programs do not always result in faster execution times!

//goroutines  are  connected with the channels , think like  channel is a pipe which connect's goroutines.

package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	fmt.Println("Doing  task with ID: ", id)
// }

func main() {

	for i := 0; i <= 10; i++ {

// go task(i)

// we can also do inline  functioning like for ex-
// now  it  perform like closure , func get's the  " i " & to run it concurrently , just write go before function declaration.
//as a best practice , always  receive the " i " , and pass into function invocation

// here  is  a  goroutine  is  started

	go func(i int) {
		fmt.Println(i)
	}(i)

}

// sleep method is inbuilt and used for  waiting of the thread execution
//in this  ex- waiting for 2 seconds , the order the running task is not proper , coz this  is  non blocking , this  is  comperatively faster then the  upper one  execution.

time.Sleep(time.Second * 2)
}

// Example of multiple Goroutiens :

package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabet() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	//numbers goroutine

	go numbers()

	//alphabet goroutine

	go alphabet()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println(" Main Terminated")

}
