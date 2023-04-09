/**
帮助理解切片的最大容量
*/
package main

import "fmt"

func main() {
	var slice []int
	for i := 0; i < 17; i++ {
		slice = append(slice, i)
		//fmt.Printf("%d %d %v \n", len(slice), cap(slice), slice)
	}

	var slice1 []int
	for i := 17; i < 34; i++ {
		slice1 = append(slice1, i)
		//fmt.Printf("%d %d %v \n", len(slice1), cap(slice1), slice1)
	}

	slice = append(slice, slice1...)
	fmt.Println(slice)
}
