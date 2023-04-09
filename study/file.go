package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	f, err := os.OpenFile(dir+"/a", os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}

	var (
		content string
		r       = bufio.NewReader(f)
	)
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		line := string(b)
		line = "dcba\n"

		content += line
	}
	f.Close()
	content += "abcd\n"
	fmt.Println(content)

	fmt.Println(ioutil.WriteFile(dir+"/a", []byte(content), os.ModePerm))
}
