package main

import "fmt"

type Tree struct {
	Val   interface{}
	Left  *Tree
	Right *Tree
}

func recursionTree(root *Tree) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	recursionTree(root.Left)
	recursionTree(root.Right)
}

func main() {
	t1 := &Tree{Val: 1}
	t2 := &Tree{Val: 2}
	t3 := &Tree{Val: 3}
	t4 := &Tree{Val: 4}
	t5 := &Tree{Val: 5}
	t6 := &Tree{Val: 6}
	t7 := &Tree{Val: 7}
	t8 := &Tree{Val: 8}
	t9 := &Tree{Val: 9}

	t1.Left = t2
	t1.Right = t3

	t2.Left = t4
	t2.Right = t5
	t3.Left = t6
	t3.Right = t7

	t4.Left = t8
	t5.Left = t9

	recursionTree(t1)
}
