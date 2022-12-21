package main

import "fmt"

type Student struct {
	Name   string
	Points []int
}

func main() {

	studentA := Student{
		Name:   "Student A",
		Points: []int{3, 5, 1},
	}

	studentB := studentA
	studentA.Points[2] = 10

	fmt.Println(studentB)
}
