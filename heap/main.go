package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (i IntHeap) Len() int {
	return len(i)
}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]

}

func (ih IntHeap) Swap(i, j int) {
	//TODO implement me
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(x any) {
	//TODO implement me
	*ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() any {
	old := *ih
	l := len(old)
	result := old[l-1]
	*ih = old[0 : l-1]
	return result
}

func main() {
	intHeap := &IntHeap{10, 7, 12, 3, 9}
	heap.Init(intHeap)
	heap.Push(intHeap, 4)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(intHeap))
	}
}
