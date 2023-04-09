// 小偷隔开一个房间偷最高金额
// dp
// https://leetcode.cn/problems/house-robber/
package main

import (
	"fmt"
	"my-go/algorithm/utils"
)

func main() {
	fmt.Println(rob([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}))
}

// 滚动数组，只需记录前两间房屋最大的总金额
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := utils.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, utils.Max(first+nums[i], second)
	}
	return second
}

// 动态规划数组
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dpArray := make([]int, len(nums))
	dpArray[0] = nums[0]
	dpArray[1] = utils.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dpArray[i] = utils.Max(dpArray[i-2]+nums[i], dpArray[i-1])
	}
	return dpArray[len(nums)-1]
}

// 备忘录
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	return dpByRob(nums, 0, memo)
}

func dpByRob(nums []int, k int, memo []int) int {
	n := len(nums)
	if k >= n {
		return 0
	}
	if memo[k] != -1 {
		return memo[k]
	}

	var res int
	for i := k; i < n; i++ {
		tmp := nums[i] + dpByRob(nums, i+2, memo)
		res = utils.Max(res, tmp)
	}
	memo[k] = res
	return res
}
