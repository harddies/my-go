package main

import (
	"fmt"
	"time"
)

func countDown(launch chan int) {
	t := time.Tick(time.Second)
	for cd := 10; cd > 0; cd-- {
		fmt.Println(cd)
		<-t
	}
	launch <- -1
}

func lau(l chan int) {
	t := time.Tick(time.Second)
	for cd := 5; cd > 0; cd-- {
		fmt.Println(cd)
		<-t
	}
	l <- -1
}

func main() {

	launch, l := make(chan int), make(chan int)
	go countDown(launch)
	go lau(l)
	for {
		select {
		case <-launch:
			fmt.Println("琳宝宝！！一心爱你呀！！")
		case <-l:
			fmt.Println("over")
		}
	}

}
