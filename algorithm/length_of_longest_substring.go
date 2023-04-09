// 不重复字符最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package main

import (
	"fmt"
	"my-go/algorithm/utils"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	var longest int
	m := make(map[byte]struct{})
	for left, right := 0, 0; right < len(s); {
		_, ok := m[s[right]]
		if !ok {
			// 用right - left + 1来现场计算长度，避免再开辟一个变量来做临时存储，并且新开辟一个临时变量存储会引发更多位置问题
			longest = utils.Max(longest, right-left+1)
			m[s[right]] = struct{}{}
			right++
		} else {
			delete(m, s[left])
			left++
		}
	}
	return longest
}
