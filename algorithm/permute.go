// 排列
// 回溯
// https://leetcode.cn/problems/permutations/
package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var res [][]int
	left := make([]int, len(nums))
	copy(left, nums)
	recursionByPermute(nums, []int{}, left, &res)
	return res
}

func recursionByPermute(nums, unit, left []int, res *[][]int) {
	count := len(unit)
	n := len(nums)
	if count == n {
		newUnit := make([]int, count)
		copy(newUnit, unit)
		*res = append(*res, newUnit)
		return
	}

	for i := 0; i < len(left); i++ {
		unit = append(unit, left[i])

		newLeft := make([]int, len(left))
		copy(newLeft, left)
		newLeft = append(newLeft[:i], newLeft[i+1:]...)
		recursionByPermute(nums, unit, newLeft, res)
		unit = unit[:len(unit)-1]
	}
}
