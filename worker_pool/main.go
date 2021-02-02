// start worker pool of 3 member
// read jobs from the same channel int
// return result channel of int * 2

// text output : worker id started job ; worker id finished job ;
// problem 1 of this method is using a large channel of jobs with consume high ram
// problem 2 feeding channel all the value before process it
package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	fmt.Println(len(jobs))

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	close(jobs)
	for j := 1; j <= numJobs; j++ {
		<-results
	}

}

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
