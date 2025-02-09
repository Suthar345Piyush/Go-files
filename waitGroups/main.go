//wait groups  in golang  :
//Waitgroups are used for executing/completion of collection of Goroutiens.
//The "wait" method is  used to block the execution of the program until all the goroutines have called  "Done".
//Ensuring synchronization among the goroutines.
//wait group has three things : 1. Add() , 2. Done() , 3. Wait()

//they are used with pointers

package main

import (
	"fmt"
	"sync"
)

//the waitgroup is refferd in task function as " w " , and in main function as " wg " , but they are same , just different variable name.

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
