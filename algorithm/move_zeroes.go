// 移动0到数组最后
// https://leetcode.cn/problems/move-zeroes/
package main

import "fmt"

func main() {
	moveZeroes1([]int{0, 1, 0, 3, 12})
}

// 优雅
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// 非优雅
func moveZeroes1(nums []int) {
	for i, j := 0, len(nums); i < j; {
		if nums[i] != 0 {
			i++
		} else {
			moveZero(nums[i:j])
			j--
		}
	}
	fmt.Println(nums)
}

func moveZero(ns []int) {
	for i := 0; i < len(ns)-1; i++ {
		if i+1 <= len(ns)-1 {
			ns[i] = ns[i+1]
		}
	}
	ns[len(ns)-1] = 0
}
