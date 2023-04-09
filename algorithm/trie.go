// 前缀树数据结构
package main

import (
	"fmt"
	"sort"
)

var words = []*Word{
	{
		S: "我想吃泡面",
		W: 40,
	},
	{
		S: "我想打游戏",
		W: 50,
	},
	{
		S: "我想玩电脑",
		W: 70,
	},
	{
		S: "我想吃香蕉",
		W: 60,
	},
	{
		S: "我打怪兽",
		W: 100,
	},
}

type Node struct {
	Pass   int
	End    bool
	Weight int
	Paths  map[rune]*Node
}

func (n *Node) FindWord(word string) []*Word {
	if n == nil {
		return nil
	}

	runes := []rune(word)
	node := n
	for _, r := range runes {
		node = node.Find(r)
	}
	return node.Iterator(word)
}

func (n *Node) Find(r rune) *Node {
	if n.End {
		return n
	}
	return n.Paths[r]
}

func (n *Node) Iterator(findWord string) []*Word {
	if n == nil {
		return nil
	}

	var res []*Word
	if n.End {
		res = append(res, &Word{
			S: findWord,
			W: n.Weight,
		})
	}

	for r, node := range n.Paths {
		res = append(res, node.Iterator(findWord+string(r))...)
	}
	return res
}

func (n *Node) Insert(r rune) *Node {
	n.Pass++
	if _, ok := n.Paths[r]; !ok {
		n.Paths[r] = NewNode()
	}
	return n.Paths[r]
}

func (n *Node) InsertWord(word *Word) {
	n.Pass++
	tmp := []rune(word.S)
	node := n
	for _, t := range tmp {
		node = node.Insert(t)
	}
	node.End = true
	node.Weight = word.W
}

func NewNode() *Node {
	node := &Node{
		Pass:   0,
		End:    false,
		Weight: 0,
		Paths:  make(map[rune]*Node),
	}
	return node
}

type Word struct {
	S string
	W int
}

type Words []*Word

func (w Words) Len() int {
	return len(w)
}

func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w Words) Less(i, j int) bool {
	return w[i].W >= w[j].W
}

func NewTrieTree() *Node {
	root := NewNode()
	for _, word := range words {
		node := root
		node.InsertWord(word)
	}
	return root
}

func main() {
	tree := NewTrieTree()
	ws := tree.FindWord("我")
	sort.Sort(Words(ws))

	for _, word := range ws {
		fmt.Println(word.S, word.W)
	}
}
