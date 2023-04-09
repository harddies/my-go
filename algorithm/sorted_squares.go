// 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/
package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-1, 2, 2}))
}

func sortedSquares(nums []int) []int {
	count := len(nums)
	if count <= 0 {
		return []int{}
	}

	minKey := -1
	min := 10000 * 10000
	total := count - 1
	// 这里写复杂了，更好的方式是：区分负数与非负数的临界点即可，即获取到最后一个是负数的位置即可
	for i := 0; i <= total; i++ {
		square := nums[i] * nums[i]
		if square <= min {
			min = square
			minKey = i
		}
	}
	fmt.Println(minKey)

	left := minKey
	right := minKey
	var res []int
	for left >= 0 || right <= count-1 {
		if left < 0 {
			res = append(res, nums[right]*nums[right])
		} else if right > total {
			res = append(res, nums[left]*nums[left])
		} else if left == right {
			res = append(res, nums[left]*nums[left])
		} else {
			mins := nums[left] * nums[left]
			maxs := nums[right] * nums[right]
			if mins <= maxs {
				res = append(res, mins)
				left--
			} else {
				res = append(res, maxs)
				right++
			}
			continue
		}

		left--
		right++
	}
	return res
}

// 两头双指针，最大值往新数组的最后开始填充
func sortedSquares1(nums []int) []int {
	count := len(nums)
	total := count - 1
	res := make([]int, count)
	left := 0
	right := total

	i := total
	for left <= right {
		if nums[right]*nums[right] >= nums[left]*nums[left] {
			res[i] = nums[right] * nums[right]
			right--
		} else {
			res[i] = nums[left] * nums[left]
			left++
		}
		i--
	}
	return res
}
