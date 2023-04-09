/*
由于多个goroutine并发读写同一个map，会报出错误：fatal error: concurrent map iteration and map write
所以可以通过建立struct，数据字段，锁字段，通过锁字段去做读写锁，来保证多goroutine并发读写同一个map报错的问题
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type MyData struct {
	Data map[int]interface{}
	lock sync.RWMutex
}

func main() {
	var d = &MyData{}
	d.Data = make(map[int]interface{})

	for i := 0; i < 1000; i++ {
		go func(i int) {
			d.lock.Lock()
			d.Data[i] = i
			d.lock.Unlock()
		}(i)
	}

	go func() {
		d.lock.Lock()
		d.Data[10000] = 10000
		d.lock.Unlock()
	}()

	go func() {
		d.lock.Lock()
		fmt.Println(len(d.Data))
		d.lock.Unlock()
	}()

	time.Sleep(1 * time.Second)
}
