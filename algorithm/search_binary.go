// 二分查找
// binarySearch 采用双指针
// search 采用递归
// https://leetcode.cn/problems/binary-search/
package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{9}, 9))
	fmt.Println(search([]int{2, 5}, 5))
}

func search(nums []int, target int) int {
	count := len(nums)
	// base case 1
	if count <= 0 {
		return -1
	}

	end := len(nums) - 1
	mid := (0 + end) / 2
	// base case 2
	if end == 0 {
		if nums[mid] == target {
			return mid
		} else {
			return -1
		}
	}
	left := div(nums, 0, mid, target)
	right := div(nums, mid+1, end, target)
	if left != -1 {
		return left
	}
	if right != -1 {
		return right
	}
	return -1
}

func div(nums []int, start, end, target int) int {
	mid := (end + start) / 2
	if nums[mid] == target {
		return mid
	}

	if target >= nums[start] && target <= nums[mid] {
		return div(nums, start, mid, target)
	}

	if target >= nums[mid] && target <= nums[end] {
		return div(nums, mid+1, end, target)
	}

	return -1
}

func binarySearch(nums []int, target int) int {
	count := len(nums)
	if count <= 0 {
		return -1
	}

	left := 0
	right := count - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if left == right {
			break
		}
		if target > nums[mid] {
			// 中间后一位
			left = mid + 1
		} else if target < nums[mid] {
			// 中间前一位
			right = mid - 1
		}
	}
	return -1
}
