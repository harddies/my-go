package main

import (
	"github.com/qianlnk/pgbar"
	"time"
)

func main() {
	pgb := pgbar.New("test")
	b := pgb.NewBar("123", 25)
	for i := 0; i < 25; i++ {
		b.Add()
		time.Sleep(time.Millisecond)
	}
}
