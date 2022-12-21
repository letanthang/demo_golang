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
	fmt.Printf("%v %d %d\n", studentA.Points, len(studentA.Points), cap(studentA.Points))

	studentB := studentA
	studentA.Points = append(studentA.Points, 7)
	studentB.Points[2] = 10

	fmt.Println(studentB) // [3,5,10]
}
