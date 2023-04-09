package main

var s *int
// go build -gcflags '-m -l' variable_escape_analysis.go
func main() {
	m := 3
	call(&m)
	m = m + *s
}

func call(i *int) {
	*i = 4
}

/*
package main
import "fmt"

func foo() *int {
	t := 3

	return &t
}
func main() {
	x := foo()
	fmt.Println(*x)
}*/