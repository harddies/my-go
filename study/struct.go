package main

import "fmt"

type A struct {
	Test string
}

type B struct {
	A
	Code string
}

func (a *A) Set() {
	a.Test = "TEST"
}

func main() {
	b := new(B)
	b.Test = "1234"
	fmt.Println(b)
	b.Set()
	fmt.Println(b)
}
