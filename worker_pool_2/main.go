// start worker pool of 3 member
// read jobs from the same channel int
// return result channel of int * 2
// text output : worker id started job ; worker id finished job ;

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	const numWorkers = 1000
	const numJobs = 1000
	jobs := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for job := range jobs {
				process(i, job)
			}
		}(i)
	}
	// feeding channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("elapsed time", elapsed)

}

func process(id int, job int) {
	fmt.Println("worker", id, "started job", job)
	time.Sleep(time.Second)
	fmt.Println("worker", id, "finished job", job)
}
