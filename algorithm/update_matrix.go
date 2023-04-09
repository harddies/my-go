// 矩阵
// 不易懂，可多看看
// https://leetcode.cn/problems/01-matrix/
package main

import (
	"fmt"
	"math"
	"my-go/algorithm/utils"
)

func main() {
	fmt.Println(dp([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
	fmt.Println(bfs([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}

// 动态规划
func dp(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = math.MaxInt32
			if mat[i][j] == 0 {
				res[i][j] = 0
			}
		}
	}

	// 水平向左 和 竖直向上找，顺序遍历，往回找
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 水平向左
			if i-1 >= 0 {
				res[i][j] = utils.Min(res[i][j], res[i-1][j]+1)
			}
			// 竖直向上
			if j-1 >= 0 {
				res[i][j] = utils.Min(res[i][j], res[i][j-1]+1)
			}
		}
	}

	// 水平向右 和 竖直向下找，倒序遍历，往前找
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 水平向右
			if i+1 < m {
				res[i][j] = utils.Min(res[i][j], res[i+1][j]+1)
			}
			// 竖直向下
			if j+1 < n {
				res[i][j] = utils.Min(res[i][j], res[i][j+1]+1)
			}
		}
	}

	return res
}

// 广度优先遍历
// 因为是优先从0开始入队，先处理0周围的1，再处理1周围的1，以此类推，即由0扩散搜索，把最近搜索到的不为0的，写为最近距离值，并标记为已搜索后，就不会再有后续入队的更远距离的对齐的覆盖
// 又因为 res[ni][nj] = res[i][j] + 1 又是类似范围层级的概念，所以当 res[ni][nj] = res[i][j] + 1, s[ni][nj] = true 必然是更远范围未被搜索的会去入队
// 入队肯定是先处理0->0->0->1->1->2->2...，而不是 0->1->2->0->1->2...
func bfs(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	res := make([][]int, m)

	q := make([][2]int, 0, m*n)
	s := make([][]bool, m)

	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		s[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, [2]int{i, j})
				s[i][j] = true
			}
		}
	}

	schema := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		i, j := v[0], v[1]
		for _, sch := range schema {
			ni, nj := i+sch[0], j+sch[1]
			if ni >= 0 && nj >= 0 && ni < m && nj < n && !s[ni][nj] {
				res[ni][nj] = res[i][j] + 1
				q = append(q, [2]int{ni, nj})
				s[ni][nj] = true
			}
		}
	}

	return res
}
