package main

import "fmt"

var ch = make(chan int)

func main() {
	for i := 1; i <= 5; i++ {
		go worker(i)
	}

	for i := 1; i <= 50; i++ {
		<-ch
	}
}

func worker(i int) {
	for j := 1; j <= 10; j++ {
		go subWorker(i, j)
	}
}

func subWorker(i, j int) {
	fmt.Printf("父协程ID：%d，子协程ID：%d\n", i, j)
	ch <- i * j
}
