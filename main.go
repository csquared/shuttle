package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 20, "number of lines to buffer before shipping")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("Must include http address as first argument")
	}
	reader := io.TeeReader(os.Stdin, os.Stdout)
	url := flag.Args()[0]
	fmt.Println("url:", url, "n:", n)
	pipeToHttp(reader, url, n)
}
