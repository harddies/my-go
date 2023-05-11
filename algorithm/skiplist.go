package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	sl := NewSkipList()
	for i := 1; i <= 20; i++ {
		sl.Insert(i, strconv.Itoa(i))
	}

	sl.Insert(4, "4")
	fmt.Println(sl.Search(21).GetValue())
	sl.Output()
}

type SkipNode struct {
	key   int
	value interface{}
	right *SkipNode
	down  *SkipNode
}

func (sn *SkipNode) GetValue() interface{} {
	if sn != nil {
		return sn.value
	}
	return nil
}

func newSkipNode(key int, value interface{}) *SkipNode {
	return &SkipNode{
		key:   key,
		value: value,
	}
}

type SkipList struct {
	head     *SkipNode
	curLevel int
	random   *rand.Rand
}

func (sl *SkipList) Search(key int) *SkipNode {
	n := sl.head
	for n != nil {
		switch {
		case n.key == key:
			return n
		case n.right == nil:
			n = n.down
		case n.right.key > key:
			n = n.down
		default:
			n = n.right
		}
	}
	return nil
}

func (sl *SkipList) Remove(key int) {
	n := sl.head
	for n != nil {
		switch {
		case n.right == nil:
			n = n.down
		case n.right.key == key:
			n.right = n.right.right
			n = n.down
		case n.right.key > key:
			n = n.down
		default:
			n = n.right
		}
	}
}

func (sl *SkipList) Insert(key int, value interface{}) {
	const MaxLevel = 32
	var n *SkipNode
	/* // 插入时不可重复的逻辑，同key覆盖
	n = sl.Search(key)
	if n != nil {
		n.value = value
		return
	}*/

	var stack []*SkipNode
	n = sl.head
	for n != nil {
		switch {
		case n.right == nil:
			stack = append(stack, n)
			n = n.down
		case n.right.key > key:
			stack = append(stack, n)
			n = n.down
		default:
			n = n.right
		}
	}

	level := 1
	var downNode *SkipNode

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		newNode := newSkipNode(key, value)
		newNode.down = downNode
		downNode = newNode
		if node.right == nil {
			node.right = newNode
		} else {
			newNode.right = node.right
			node.right = newNode
		}

		if level > MaxLevel {
			break
		}
		if sl.random.Intn(10) > 5 {
			break
		}
		level++

		if level > sl.curLevel {
			sl.curLevel = level
			newHeadNode := newSkipNode(0, nil)
			newHeadNode.down = sl.head
			sl.head = newHeadNode
			stack = append(stack, newHeadNode)
		}
	}
}

func (sl *SkipList) Output() {
	n := sl.head
	index := 1
	last := n
	for last.down != nil {
		last = last.down
	}

	for n != nil {
		enumNode := n.right
		enumLast := last.right
		fmt.Printf("%-8s", "head->")

		for enumNode != nil && enumLast != nil {
			if enumLast.key == enumNode.key {
				fmt.Printf("%-6s", fmt.Sprintf("%d", enumLast.key)+"->")
				enumLast = enumLast.right
				enumNode = enumNode.right
			} else {
				enumLast = enumLast.right
				fmt.Printf("%-6s", "")
			}
		}

		n = n.down
		index++
		fmt.Println()
	}
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:     newSkipNode(0, nil),
		curLevel: 0,
		random:   rand.New(rand.NewSource(time.Now().Unix())),
	}
}
