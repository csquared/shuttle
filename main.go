package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Must include http address as first argument")
	}
	reader := io.TeeReader(os.Stdin, os.Stdout)
	pipeToHttp(reader, args[0], 20)
}
