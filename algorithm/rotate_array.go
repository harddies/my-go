// 轮转数组
// https://leetcode.cn/problems/rotate-array/
package main

import "fmt"

func main() {
	rotate2([]int{1, 2, 3, 4, 5, 6}, 3)
}

// 开辟新数组来copy回去
func rotate(nums []int, k int) {
	var newNums = make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 硬翻转
func rotate1(nums []int, k int) {
	count := len(nums)
	left := (count - 1) / 2
	right := (count-1)/2 + (1 - count%2)
	for left >= 0 && right <= count-1 {
		nums[left], nums[right] = nums[right], nums[left]

		left--
		right++
	}
	fmt.Println(nums)

	k %= count
	count = len(nums[:k])
	left = (count - 1) / 2
	right = (count-1)/2 + (1 - count%2)
	for left >= 0 && right <= count-1 {
		nums[left], nums[right] = nums[right], nums[left]

		left--
		right++
	}

	count = len(nums)
	left = (k + count - 1) / 2
	right = (k+count-1)/2 + (1 - len(nums[k:])%2)
	for left >= k && right <= count-1 {
		nums[left], nums[right] = nums[right], nums[left]

		left--
		right++
	}
	fmt.Println(nums)
}

// 通过取余改变为合法的k下标来翻转
func revert(ns []int) {
	for i, count := 0, len(ns); i < count/2; i++ {
		ns[i], ns[count-1-i] = ns[count-1-i], ns[i]
	}
}

func rotate2(nums []int, k int) {
	revert(nums)
	k %= len(nums)
	revert(nums[:k])
	revert(nums[k:])
	fmt.Println(nums)
}
