package main

import (
	"fmt"
	"my-go/study/sort"
)

func main() {
	arr := []int{3, 8, 1, 10, 2, 5, 4}

	//fmt.Println(sort.QuickSort(arr))

	//fmt.Println(arr)

	fmt.Println(sort.PushUnshiftSort(arr))

	//fmt.Println(arr)
}
