// 搜索岛屿数量
// https://leetcode.cn/classic/problems/number-of-islands/description/
package main

import "fmt"

func main() {
	grid := [][]byte{
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}

	var count int
	for i := 0; i < m; i++ {
		n := len(grid[i])
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
				grid[i][j] = 0
				deepSearch(grid, i-1, j)
				deepSearch(grid, i+1, j)
				deepSearch(grid, i, j-1)
				deepSearch(grid, i, j+1)
			}
		}
	}
	return count
}

func deepSearch(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
		return
	}

	if grid[i][j] == 1 {
		grid[i][j] = 0
		deepSearch(grid, i-1, j)
		deepSearch(grid, i+1, j)
		deepSearch(grid, i, j-1)
		deepSearch(grid, i, j+1)
	}
}
