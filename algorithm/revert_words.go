// 反转单词3
// https://leetcode.cn/problems/reverse-words-in-a-string-iii/
package main

import "fmt"

func main() {
	fmt.Println(reverseWords("I love u"))
}

func reverseWords(s string) string {
	i := 0
	sb := []byte(s)
	n := len(sb)
	for i < n {
		start := i
		for i < n && sb[i] != ' ' {
			i++
		}

		left, right := start, i-1
		for left < right {
			sb[left], sb[right] = sb[right], sb[left]
			left++
			right--
		}

		for i < n && sb[i] == ' ' {
			i++
		}
	}
	return string(sb)
}

func reverseWords1(s string) string {
	var swapping bool
	var mark int
	var sb = []byte(s)
	for left, right := 0, 0; right <= len(sb); {
		if left >= right && swapping {
			swapping = false
			left, right = mark+2, mark+2
			continue
		}
		if (right+1 < len(sb) && sb[right+1] == ' ') || (right == len(sb)-1 && sb[right] != ' ') || swapping {
			if !swapping {
				mark = right
			}
			sb[left], sb[right] = sb[right], sb[left]
			left++
			right--
			swapping = true
			continue
		}
		right++
	}
	return string(sb)
}
