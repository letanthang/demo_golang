package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Task1(&wg)
	go Task2(&wg)
	Task3()
	wg.Wait()

}

func Task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("task1 done")

}

func Task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("task2 done")

}

func Task3() {
	fmt.Println("task3 done")
}
