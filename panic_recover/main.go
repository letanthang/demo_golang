package main

import (
	"fmt"
	"log"
)

func safelyDo(work int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	fmt.Println(100 / work)
}

func main() {
	safelyDo(0)
	fmt.Println("Program end")
}
