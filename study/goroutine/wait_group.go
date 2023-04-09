package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	t := time.Tick(time.Second)

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)

		<-t
	}

	wg.Wait()
	fmt.Println("over")
}
