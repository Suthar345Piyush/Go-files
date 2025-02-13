//channels , that we talk about earlier are unbuffered channels , in which sends and receives are blocking

//Buffered Channels = This means , send to a buffered channel , only block when buffer is full , and receive from buffer channels , only block when the buffer is empty

//we have to add some capacity while making the channel with make keyword/function

//example

// package main

// import "fmt"

// func main() {

//here 2 is capacity,means this channels now only receive 2 , without being bloked

// 	ch := make(chan string, 2)
// 	ch <- "piyush"
// 	ch <- "suthar"
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

//another complex example

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("Successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int, 2)
// 	go write(ch)

// 	time.Sleep(2 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 		time.Sleep(2 * time.Second)
// 	}
// }

//DEADLOCK in buffered channels

// package main

// import "fmt"

// func main() {
// 	ch := make(chan string, 2)
// 	ch <- "piyush"
// 	ch <- "mehnat kar"

//here buffer is  full , reached its capacity
// 	ch <- "suthar"
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)

// }

//here program will panic , becoz there is no concurrent routine  reading from this channel
//if i write three print statement , then also it will panic

//CLOSING BUFFERED CHANNELS:

//in buffered channels case , we can read data from even closed channels

// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 5)
//writing from channel
// ch <- 5
// ch <- 6
// ch <- 7
// close(ch)
//reading data from channel
//here n=5
// 	n, open := <-ch
// 	fmt.Printf("Recieved: %d , open: %t\n", n, open)
// 	n, open = <-ch
// 	fmt.Printf("Recieved: %d , open: %t\n", n, open)
// 	n, open = <-ch
// 	fmt.Printf("recieved: %d , open: %t\n", n, open)
// }

//this can be written using for range loop like this

// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 5)
// 	ch <- 5
// 	ch <- 6
// 	close(ch)

// 	for n := range ch {
// 		fmt.Println("received: ", n)
// 	}
// }

//LENGTH vs CAPACITY

