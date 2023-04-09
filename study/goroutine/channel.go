package main

import (
	"fmt"
)

var count = 0

func cCount(ch chan int) {
	ch <- count
	fmt.Println("count: ", count)
	count++
}

func main() {

	chs := make([]chan int, 11)

	for i := 0; i <= 10; i++ {
		chs[i] = make(chan int)
		go cCount(chs[i])
	}

	var value int
	for _, ch := range chs {
		value = <-ch
		if value < 10 {
			fmt.Println(value)
		}
	}
}

/*func main() {
	f, err := os.OpenFile("/go.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	chs := make([] chan bool, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan bool)
		go func(i int) {
			f.WriteString(strconv.Itoa(i) + "\n")
			chs[i] <- true
		}(i)
	}

	for _, ch := range chs {
		<-ch
	}
}*/
