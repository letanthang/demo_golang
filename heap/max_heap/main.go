package main

import (
	"container/heap"
	"fmt"
)

type MaxIntHeap []int

func (m MaxIntHeap) Len() int {
	return len(m)
}

func (m MaxIntHeap) Less(i, j int) bool {
	// max use ">", minHeap use <
	return m[i] > m[j]
}
func (m MaxIntHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxIntHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxIntHeap) Pop() any {
	old := *m
	l := len(old)
	result := old[l-1]
	*m = old[0 : l-1]
	fmt.Println(l-1, old, m)
	return result
}

func main() {
	maxIntHeap := &MaxIntHeap{1, 9, 3, 7}
	heap.Init(maxIntHeap)
	heap.Push(maxIntHeap, 4)
	fmt.Println(maxIntHeap)
	fmt.Println("maximum", heap.Pop(maxIntHeap))
	for maxIntHeap.Len() > 0 {
		v := heap.Pop(maxIntHeap)
		fmt.Println(v)
	}
}
