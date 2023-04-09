/**
帮助理解无序集合map，以及range的使用
*/
package main

import "fmt"

func main() {
	var m = map[string]string{"胖子": "老李的孙子", "老陈": "老李的儿子"}
	for name, nick := range m {
		fmt.Println(name + "是" + nick)
	}

	delete(m, "胖子")

	nick, ok := m["胖子"]
	if ok {
		fmt.Println(nick)
	} else {
		fmt.Println("不存在")
	}

	fmt.Println(m)

	m["胖子2"] = "老李的孙子2"
	fmt.Println(m)

	m["胖子1"] = "老李的孙子1"
	fmt.Println(m)
}
