package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("brew", "update")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("command.Execute cmd.StdoutPipe error, ", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln("command.Execute cmd.StderrPipe error, ", err)
		return
	}

	go output(stdout)
	go output(stderr)

	if err = cmd.Start(); err != nil { // 运行命令
		log.Fatalln("command.Execute cmd.Start error, ", err)
		return
	}
	err = cmd.Wait()
}

var output = func(pipe io.ReadCloser) {
	log.Println("command output start")
	var br = bufio.NewReader(pipe)
	for {
		line, _, e := br.ReadLine()
		if e != nil {
			log.Fatalln("fail to read line, err=", e)
			break
		}
		log.Println(string(line))
	}
	log.Println("command output finish")
}
