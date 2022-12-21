package main

import "fmt"

type Student struct {
	Name   string
	Points []int
}

func main() {

	points := make([]int, 0, 5)
	fmt.Printf("%v %d %d\n", points, len(points), cap(points))

	//TH1: still same underlie array
	// points = append(points, 3, 5, 1)

	// TH2 new underlie array
	points = append(points, 3, 5, 1, 4, 8)

	fmt.Printf("%v %d %d\n", points, len(points), cap(points))

	studentA := Student{
		Name:   "Student A",
		Points: points,
	}
	fmt.Printf("%v %d %d\n", studentA.Points, len(studentA.Points), cap(studentA.Points))

	studentB := studentA
	studentA.Points = append(studentA.Points, 7)
	fmt.Printf("%v %d %d\n", studentA.Points, len(studentA.Points), cap(studentA.Points))
	studentB.Points[2] = 10

	fmt.Printf("%v %d %d\n", studentB.Points, len(studentB.Points), cap(studentB.Points))
	fmt.Printf("%v %d %d\n", studentA.Points, len(studentA.Points), cap(studentA.Points))
	// fmt.Println(studentB) // [3,5,10]
}
