// 限定背包重量装满最大价值物品
// 贪心算法，每次处理小问题的时候都选最优方法
package main

import "fmt"

func main() {
	fmt.Println(bagValue(150, []float64{35, 30, 60, 50, 40, 10, 25}, []float64{10, 40, 30, 50, 35, 40, 30}))
}

func bagValue(weight float64, thingWeights []float64, thingValues []float64) float64 {
	n := len(thingWeights)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if thingValues[i]/thingWeights[i] < thingValues[j]/thingWeights[j] {
				thingValues[i], thingValues[j] = thingValues[j], thingValues[i]
				thingWeights[i], thingWeights[j] = thingWeights[j], thingWeights[i]
			}
		}
	}
	fmt.Println(thingWeights, thingValues)
	var v float64
	for i := 0; i < n; i++ {
		if weight < thingWeights[i] {
			continue
		}
		weight -= thingWeights[i]
		v += thingValues[i]
	}
	return v
}
