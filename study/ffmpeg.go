package main

import (
	"fmt"
	"log"

	"github.com/giorgisio/goav/avformat"
)

var inputFileName = "/Users/ronnie/Desktop/test.mp3"

func main() {
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	if avformat.AvformatOpenInput(&ctx, inputFileName, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	fmt.Printf("bit rate: %+v\n", ctx.BitRate())
}
