package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3}
	b := []int{2, 7}
	fmt.Println(findMedian(a, b))
}

func findMedian(a, b []int) float64 {
	aLen := len(a)
	bLen := len(b)
	var i, j = 0, 0
	var result []int
	if aLen <= 0 && bLen <= 0 {
		return 0
	}
	if aLen <= 0 && bLen > 0 {
		result = b
	} else if aLen > 0 && bLen <= 0 {
		result = a
	} else {
		for i < aLen && j < bLen {
			if a[i] <= b[j] {
				result = append(result, a[i])
				i++
			} else {
				result = append(result, b[j])
				j++
			}

			if i >= aLen {
				result = append(result, b[j:]...)
				break
			}
			if j >= bLen {
				result = append(result, a[i:]...)
				break
			}
		}
	}
	fmt.Println(result)

	rLen := len(result)
	if len(result) == 0 {
		return 0
	}
	if rLen%2 == 1 {
		return float64(result[(rLen / 2)])
	} else {
		return (float64(result[rLen/2-1]) + float64(result[rLen/2])) / 2
	}
}
