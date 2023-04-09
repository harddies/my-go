/**
函数挂载到结构体的方法
*/
package main

import "fmt"

type Test struct {
	name  string
	level string
	do    string
	whom  string
}

func main() {
	var t Test

	t.name = "李逸昕"
	t.level = "超级"
	t.do = "爱"
	t.whom = "曾琳"

	t.combine()
}

func (t Test) combine() {
	fmt.Printf("%s%s%s%s\n", t.name, t.level, t.do, t.whom)
}
