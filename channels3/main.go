//closing of channels , using close method comes with the v , ok (value and ok for surity)

package main

import "fmt"

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

//it will give ok or false for the channel (ye close)

func main() {
	ch := make(chan int)
	go producer(ch)

	//we can write the for range loop for the values

	for v := range ch {
		fmt.Println("recevied", v)
	}

	// for {
	// 	v, ok := <-ch
	// 	// if ok == false {
	// 	// 	break
	// 	// }
	// 	fmt.Println("received", v, ok)
	// }
}

// //

//More optimized code of the calcsquare / cubes using for range loop

package main

import "fmt"

//seperate digits function
func digits(number int, digitchnl chan int) {
	for number != 0 {
		digit := number % 10
		digitchnl <- digit
		number /= 10
	}

	close(digitchnl)
}

//square function
//using digits function goroutine

func calcSquare(number int, squareop chan int) {
	sum := 0
	dgch := make(chan int)
	go digits(number, dgch)

	for digit := range dgch {
		sum += digit * digit
	}

	squareop <- sum

}

//cube function , also using digits function goroutine

func calcCube(number int, cubeop chan int) {
	sum := 0
	dgch := make(chan int)
	go digits(number, dgch)

	for digit := range dgch {
		sum += digit * digit * digit
	}

	cubeop <- sum
}

func main() {
	number := 50
	sqrch := make(chan int)
	cbch := make(chan int)

	go calcSquare(number, sqrch)
	go calcCube(number, cbch)

	squares, cubes := <-sqrch, <-cbch

	fmt.Println("final ans is: ", squares+cubes)
}
