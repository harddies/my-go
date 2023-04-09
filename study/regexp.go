package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `④ 30万≤X＜50万 3000 元/月
⑤ 50万≤X＜70万 6000 元/月
⑥ 70万≤X 9000 元/月`
	reg := regexp.MustCompile("⑥[\\n\\s]*(\\d+)[\\n\\s]*万≤X[＜\\d+万]*[\\n\\s]*(\\d+)[\\n\\s]*元/月")
	result := reg.FindStringSubmatch(str)
	fmt.Println(result, len(result))
	for _, res := range result {
		fmt.Println(res)
	}
}
