// 图像渲染
// https://leetcode.cn/problems/flood-fill/
package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0))
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	render(image, sr, sc, image[sr][sc], color)
	return image
}

func render(image [][]int, sr, sc, oldColor, color int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldColor || oldColor == color {
		return
	}
	image[sr][sc] = color
	render(image, sr-1, sc, oldColor, color)
	render(image, sr+1, sc, oldColor, color)
	render(image, sr, sc-1, oldColor, color)
	render(image, sr, sc+1, oldColor, color)
}
