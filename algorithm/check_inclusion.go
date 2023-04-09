// 字符串的排列，s2是否包含s1的排列
// https://leetcode.cn/problems/permutation-in-string/
package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("bca", "abcabcbb"))
}

// 利用数组
func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt1, cnt2 [26]int
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

// 粗暴map记录
func checkInclusion1(s1 string, s2 string) bool {
	s1len := len(s1)
	sb1 := []byte(s1)
	m := make(map[byte]int)
	for i := 0; i < s1len; i++ {
		m[sb1[i]]++
	}

	for i, j := 0, s1len-1; j < len(s2); {
		tm := make(map[byte]int)
		for x := i; x <= j; x++ {
			tm[s2[x]]++
		}
		var inclusion = true
		for x := range tm {
			if tm[x] != m[x] {
				inclusion = false
			}
		}
		if inclusion {
			return true
		}
		i++
		j++
	}
	return false
}
