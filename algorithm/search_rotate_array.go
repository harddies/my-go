// 旋转数组搜索
// https://leetcode.cn/problems/search-in-rotated-sorted-array/
package main

import "fmt"

func main() {
	fmt.Println(searchRotateArray([]int{5, 6, 7, 1, 2, 3, 4}, 7))
}

func searchRotateArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[0] <= nums[mid] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
