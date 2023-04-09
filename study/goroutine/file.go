package main

import (
	"log"
	"os"
	"strconv"
)

var fileName = "/go.log"

func main() {

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		f.WriteString(strconv.Itoa(i*10) + "\n")
	}

}
