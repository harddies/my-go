// 爬楼梯
// dp
// https://leetcode.cn/problems/climbing-stairs/
package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(5))
}

func climbStairs1(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func climbStairs2(n int) int {
	dpArray := make([]int, n+1)
	dpArray[0] = 1
	dpArray[1] = 1
	for i := 2; i <= n; i++ {
		dpArray[i] = dpArray[i-1] + dpArray[i-2]
	}
	return dpArray[n]
}

func climbStairs(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return dpByClimbStairs(n, memo)
}

func dpByClimbStairs(n int, memo []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if memo[n] != -1 {
		return memo[n]
	}

	var res int
	for _, v := range []int{1, 2} {
		res += dpByClimbStairs(n-v, memo)
	}
	memo[n] = res
	return res
}
