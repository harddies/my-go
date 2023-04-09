package main

import (
	"fmt"
	"time"
)

var data = make(map[int]int64)
var stopChan = make(map[int]chan bool)

func main() {
	tick := time.Tick(15 * time.Second)
	for {
		now := time.Now().Unix()
		fmt.Printf("----------------now: %+v----------------------\n", now)
		for i := 1; i <= 1; i++ {
			t := int64(i*30) + now
			fmt.Printf("i: %d, t: %d\n", i, t)
			d, ok := data[i]
			if !ok {
				data[i] = t
				stopChan[i] = trigger(i, t)
				continue
			}
			if d != t {
				fmt.Printf("data[%d]: %d\n", i, t)
				stopChan[i] <- true
				data[i] = t
				fmt.Printf("data[%d]: %d\n", i, t)
				stopChan[i] = trigger(i, t)
			}
		}
		<-tick
	}
}

func trigger(i int, t int64) chan bool {
	var (
		deltaT = t - time.Now().Unix()
		ticker = time.NewTicker(time.Duration(deltaT) * time.Second)
		isStop bool
		stopCh = make(chan bool)
	)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("exe task id: %d, t: %d\n", i, t)
				return
			case isStop = <-stopCh:
				if isStop {
					fmt.Printf("id: %d, t: %d stop\n", i, t)
					return
				}
			}
			fmt.Printf("----------------------------")
		}
	}()

	return stopCh
}
