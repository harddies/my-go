package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(324252846381864709 / 50986907)
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Sort(sort.IntSlice(candidates))
	memo := make(map[int]struct{})
	recursion(&res, candidates, []int{}, target, 0, memo)
	fmt.Println(memo)
	return res
}

func recursion(res *[][]int, candidates, path []int, target, k int, memo map[int]struct{}) {
	if target < 0 {
		return
	}
	if target == 0 {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*res = append(*res, newPath)
		return
	}

	for i := k; i < len(candidates); i++ {
		if _, ok := memo[candidates[i]]; ok {
			continue
		}
		path = append(path, candidates[i])
		memo[candidates[i]] = struct{}{}
		newMemo := make(map[int]struct{})
		recursion(res, candidates, path, target-candidates[i], i+1, newMemo)
		path = path[:len(path)-1]
	}
}
