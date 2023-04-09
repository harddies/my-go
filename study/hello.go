package main

import "fmt"

func main() {
	fmt.Println(make(map[int]int, 2))
}
func getAverage(arr []int) float32 {
	var i, sum int
	var size = len(arr)
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	fmt.Printf("sum=%d", sum)
	fmt.Printf("size=%d", size)

	avg = float32(sum) / float32(size)

	return avg
}
func SilentTime(t int32) string {
	if t > 23592359 || t < 0 {
		return ""
	}
	ts := fmt.Sprintf("%08d", t)
	fh := ts[0:2]
	fm := ts[2:4]
	th := ts[4:6]
	tm := ts[6:8]
	if fh > "23" || th > "23" {
		return ""
	}
	if fm > "59" || tm > "59" {
		return ""
	}
	return fmt.Sprintf("%s:%s-%s:%s", fh, fm, th, tm)
}
