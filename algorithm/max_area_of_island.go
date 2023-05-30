// 岛屿最大面积
// https://leetcode.cn/problems/max-area-of-island/
package main

import (
	"fmt"
	"math"
	"my-go/utils"
)

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}

	var maxArea int
	for i := 0; i < m; i++ {
		n := len(grid[i])
		for j := 0; j < n; j++ {
			var area int
			if grid[i][j] == 1 {
				area++
				grid[i][j] = 0
				area += calArea(grid, i-1, j)
				area += calArea(grid, i+1, j)
				area += calArea(grid, i, j-1)
				area += calArea(grid, i, j+1)
			}
			maxArea = utils.Max(maxArea, area)
		}
	}
	return maxArea
}

func calArea(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
		return 0
	}

	area := 1
	if grid[i][j] == 1 {
		grid[i][j] = 0
		area += calArea(grid, i-1, j)
		area += calArea(grid, i+1, j)
		area += calArea(grid, i, j-1)
		area += calArea(grid, i, j+1)
	}
	fmt.Println(math.MaxInt8)
	return area
}
