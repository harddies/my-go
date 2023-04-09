package main

import (
	"fmt"
)

func main() {
	permit := []bool{true, true, true, true}
	var val int32
	for k, v := range permit {
		val += parseVal(v, uint(k))
	}
	fmt.Println(val)
}

func parseVal(bitVal bool, bit uint) int32 {
	var val uint
	if bitVal {
		val = 1
	} else {
		val = 0
	}
	res := int32(val << bit)
	fmt.Printf("%d, %d\n", val, res)
	return res
}
