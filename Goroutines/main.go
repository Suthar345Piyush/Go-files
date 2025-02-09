//Goroutines  in Golang
//goroutines  are  light  weight  threads which is used for  multithreading & concurrency
// in simple  words -> running  task parallaly , at the same  time.
//to use them just write " go " , in the front  of  function invocation , then it  will work like a light weight thread

//this basically depands on the  core the  machine

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

		go func(i int) {
			fmt.Println(i)
		}(i)

	}

	// sleep method is inbuilt and used for  waiting of the thread execution
	//in this  ex- waiting for 2 seconds , the order the running task is not proper , coz this  is  non blocking , this  is  comperatively faster then the  upper one  execution , it  is  faster coz , they are running parallaly at the  same time

	time.Sleep(time.Second * 2)
}
