// 组合
// 回溯
// https://leetcode.cn/problems/combinations/
package main

import "fmt"

func main() {
	fmt.Println(combine(4, 3))
}

func combine(n int, k int) [][]int {
	var res [][]int
	recursionByCombine([]int{}, res, 1, n, k)
	return res
}

func recursionByCombine(unit []int, res [][]int, index, n, k int) {
	if len(unit) == k {
		newUnit := make([]int, len(unit))
		copy(newUnit, unit)
		res = append(res, newUnit)
		return
	}

	for i := index; i <= n-(k-len(unit))+1; i++ {
		unit = append(unit, i)
		recursionByCombine(unit, res, i+1, n, k)
		unit = unit[:len(unit)-1]
	}
}
