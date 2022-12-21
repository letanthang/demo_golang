package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[K constraints.Ordered, V any] struct {
	Key         K
	Value       V
	left, right *TreeNode[K, V]
}

func (n *TreeNode[K, V]) ToString() string {
	return fmt.Sprintf("TreeNode(%v,%v)", n.Key, n.Value)
}

func (n *TreeNode[K, V]) Print() {
	if n.left != nil {
		n.left.Print()
	}
	fmt.Println(n.ToString())
	if n.right != nil {
		n.right.Print()
	}
}

func (n *TreeNode[K, V]) Insert(child *TreeNode[K, V]) {
	if child.Key <= n.Key {
		if n.left == nil {
			n.left = child
		} else {
			n.left.Insert(child)
		}
	} else {
		if n.right == nil {
			n.right = child
		} else {
			n.right.Insert(child)
		}
	}
}

func main() {
	nodes := []TreeNode[int, string]{
		{Key: 1, Value: "Hihi"},
		{Key: 3, Value: "Haha"},
		{Key: 100, Value: "hyhy"},
		{Key: 50, Value: "Yeah"},
		{Key: 15, Value: "Great"},
		{Key: 30, Value: "Huhu"},
		{Key: 12, Value: "Hichic"},
	}

	head := &TreeNode[int, string]{Key: 30, Value: "Hehe"}

	for i := range nodes {
		head.Insert(&nodes[i])
	}

	head.Print()
}