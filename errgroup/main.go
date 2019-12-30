package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/sync/errgroup"
)

func main() {
	// var eg errgroup.Group
	//eg.Go(func() error)
	//eg.Wait() error

	var eg errgroup.Group

	for i := 0; i < 10; i++ {
		jobID := i
		eg.Go(func() error {
			return Job(jobID)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("Jobs error :%v\n", err)
	}

}

func Job(jobID int) error {
	if rand.Intn(10) == jobID {
		return fmt.Errorf("job %d failed", jobID)
	}

	fmt.Printf("job %d done\n", jobID)
	return nil
}
