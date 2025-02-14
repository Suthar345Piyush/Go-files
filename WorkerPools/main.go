//waitgroup -> A WaitGroup is used to wait for a collection of Goroutines to finish executing
//firtsly we have to know about how  waitgroup works before the working of the worker pools

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func process(i int, wg *sync.WaitGroup) {
// 	fmt.Println("started goroutine", i)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("goroutine %d ended\n", i)
// 	wg.Done()
// }

// func main() {
// 	no := 3
// 	var wg sync.WaitGroup
// 	for i := 0; i < no; i++ {
// 		wg.Add(1)
// 		go process(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("all go routine finished")
// }

//the buffered channels are used for impleamentation of worker pools
//like in js,nodejs these is called as the thread pool , in golang thread pool is known as the Worker Pool , working is exactly same

//Big program of worker pools

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// struct for job
type Job struct {
	id       int
	randomno int
}

// struct for result
type Result struct {
	job         Job
	sumofdigits int
}

// make  jobs and results channels

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// made a  digit function which do sum of the number individualy
func digits(number int) int {
	sum := 0
	no := number

	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(2 * time.Second)
	return sum
}

//creating worker function

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}

	wg.Done()
}

// creating the worker pool function which do the work

func createWorkerPool(noofWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noofWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(results)
}

//creating the allocation function

func allocate(noofJobs int) {
	for i := 0; i < noofJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}

	close(jobs)
}

// making the  result function using done channel
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d , input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}

	done <- true
}

func main() {
	startTime := time.Now()
	noofJobs := 100

	//goroutine of the allocate function

	go allocate(noofJobs)

	done := make(chan bool)

	//goroutine of the result function
	go result(done)

	noofWorkers := 20

	createWorkerPool(noofWorkers)
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)                             // Assign the result of endTime.Sub(startTime) to the variable diff
	fmt.Println("total time taken", diff.Seconds(), "seconds") // Use the variable diff

}
