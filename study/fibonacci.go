package main

import "fmt"

func main() {
	calc(30)
}

func calc(max int32) {
	var arr = []int32{0, 1}
	var i int32

	for i = 1; i < max; i++ {
		arr = append(arr, arr[i]+arr[i-1])
	}

	fmt.Printf("%v \n", arr)
}
