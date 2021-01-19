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
	const numWorkers = 8
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

func process2(id int, job int) {
	fmt.Println("worker", id, "started job", job)
	f := fibonacci()
	ret := 0
	for i := 0; i < 300; i++ {
		ret = f()
	}
	fmt.Println("worker", id, "finished job", job, ret)
}
func process(id int, job int) {
	start := time.Now()
	fmt.Println("worker", id, "started job", job)
	ret := finonacciRecursive(33)
	fmt.Println("worker", id, "finished job", job, ret, time.Since(start))
}

func finonacciRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return finonacciRecursive(n-1) + finonacciRecursive(n-2)
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}
