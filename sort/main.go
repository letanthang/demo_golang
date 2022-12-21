package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	is1 := []int{3, 1, 8, 2}
	is2 := make([]int, len(is1))
	copy(is2, is1)
	fmt.Println(is1, is2)
	//sort way 1
	sort.Slice(is1, func(i, j int) bool {
		return is1[i] < is1[j]
	})
	fmt.Println(is1, is2)

	//sort way 2
	sort.Sort(sort.IntSlice(is2))
	fmt.Println(is1, is2, reflect.DeepEqual(is1, is2))

	//reverse
	i := sort.Reverse(sort.IntSlice(is1))
	sort.Sort(i)
	fmt.Println(is1, is2)

	ss := []string{"Melon", "Apple", "Pinapple", "Durion"}

	sort.Sort(sort.StringSlice(ss))
	fmt.Println(ss)
}

func testArr() []int {
	return nil
}
