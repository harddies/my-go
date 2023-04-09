// 三角形路径最小之和
// dp
// https://leetcode.cn/problems/permutations/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}}))
}

func minimumTotal(triangle [][]int) int {
	memo := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		memo[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			memo[i][j] = math.MaxInt32
		}
	}
	return dpByMinimumTotal(triangle, 0, 0, memo)
}

func dpByMinimumTotal(triangle [][]int, i, j int, memo [][]int) int {
	if i >= len(triangle) || j >= len(triangle[i]) {
		return 0
	}
	if memo[i][j] != math.MaxInt32 {
		return memo[i][j]
	}

	res := math.MaxInt32
	tmpCur := triangle[i][j] + dpByMinimumTotal(triangle, i+1, j, memo)
	tmpNext := triangle[i][j] + dpByMinimumTotal(triangle, i+1, j+1, memo)
	res = min(res, tmpCur)
	res = min(res, tmpNext)
	memo[i][j] = res
	fmt.Println(i, j, res, memo)
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
