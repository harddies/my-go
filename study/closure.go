package main

import "fmt"

func main() {
	fmt.Println(test1())
}

func test1() int {
	t := 1
	return func(a int) int {
		t++
		return t * a
	}(3)
}
