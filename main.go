package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Must include http address as first argument")
	}
	pipeToHttp(os.Stdin, args[0], 20)
	fmt.Println("Done")
}
