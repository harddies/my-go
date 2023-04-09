package main

import "fmt"

func main() {
	fmt.Println(maxProfit(2, []int{1, 2, 4}))
}

func maxProfit(k int, prices []int) int {
	count := len(prices)
	if count <= 1 {
		return 0
	}

	maxes := make([]int, count-1)
	for i := 0; i < count-1; i++ {
		max := 0
		for j := i + 1; j < count; j++ {
			profit := prices[j] - prices[i]
			if profit < max {
				break
			}
			if max < profit {
				max = profit
			}
		}

		maxes[i] = max
	}

	fmt.Println(maxes)

	total := 0
	for i := 0; i < k; i++ {
		max := 0
		maxKey := -1
		for j := range maxes {
			if maxes[j] > max {
				max = maxes[j]
				maxKey = j
			}
		}
		if maxKey > -1 {
			maxes[maxKey] = 0
		}
		total += max
	}
	return total
}

/*func maxProfit1(k int, prices []int) int {
	count := len(prices)
	if count <= 1 {
		return 0
	}

	left := 0
	right := 1
	profitKeys := make([][]int, k)
	i := 0
	for {
		if left >= count || right >= count {
			break
		}

		profit := prices[right] - prices[left]
		if profit < 0 {
			left = right
			right = left+1
			continue
		}

		if profitKeys[i] == nil {
			profitKeys[i] = make([]int, 2)
		}
		if profitKeys[i][1] - profitKeys[i][0] < profit {
			profitKeys[i][0] = left
			profitKeys[i][1] = right
			right++
		}
		if right >= count {
			if profitKeys[i][1]
				i++
		}
	}
}*/
