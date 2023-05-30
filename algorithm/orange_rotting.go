// 腐烂的橘子
// https://leetcode.cn/problems/rotting-oranges/
package main

import (
	"fmt"
	"my-go/utils"
)

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	q := make([][2]int, 0, m*n)
	newGrid := make([][]int, m)
	s := make([][]bool, m)

	for i := 0; i < m; i++ {
		newGrid[i] = make([]int, n)
		s[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
				s[i][j] = true
			}
			newGrid[i][j] = grid[i][j]
		}
	}

	schema := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	minute := 0
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		i, j := v[0], v[1]
		for _, sch := range schema {
			ni, nj := i+sch[0], j+sch[1]
			if ni >= 0 && nj >= 0 && ni < m && nj < n && !s[ni][nj] {
				if grid[ni][nj] == 1 {
					q = append(q, [2]int{ni, nj})
					s[ni][nj] = true
					newGrid[ni][nj] = newGrid[i][j] + 1
					minute = utils.Max(minute, newGrid[ni][nj]-2)
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if newGrid[i][j] == 1 {
				return -1
			}
		}
	}
	return minute
}
