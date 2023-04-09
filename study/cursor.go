package main

import "fmt"

type test struct {
	value int
}

// go的函数的参数传递都是【值传递】
// 由于x是指针类型，都指向了外部指针变量指向的值，同一个内存地址
// 所以x改变了值，外部指针变量指向的值自然也会变动，这个很容易理解，和php的引用一样的概念
// 但是由于值传递，在函数里x是对外部的指针变量做了拷贝，所以对x本身的更改，不会影响到外部，只能修改值
// 所以对x=nil，是不会影响到外部指针变量的
func double(x *test)  {
	fmt.Printf("x1 %p\n", x)
	x.value += x.value
	x = nil
	fmt.Printf("x2 %p\n", x)
}

func main() {
	v := &test{
		value: 3,
	}
	fmt.Printf("v1 %p\n", v)
	double(v)

	fmt.Printf("v2 %p\n", v)
	a := v
	fmt.Printf("a %p\n", a)
	fmt.Println(v.value)
}
