package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	batchSize := 20
	lineBuffer := make([]string, batchSize)
	i := 0
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Must include http address as first argument")
	}

	var wg sync.WaitGroup
	for {
		line, err := stdin.ReadString('\n')
		// ship whatever is in the buffer at EOF
		if err == io.EOF {
			fmt.Println("EOF - exiting")
			_ = shipLinesToUrl(lineBuffer[0:i], args[0])
			break
		}
		if err != nil {
			panic("Error reading data from stdin")
		}
		// put lines in the buffer; ship them
		// when the buffer is full
		if line != "" {
			lineBuffer[i] = line
			i++
			if i == cap(lineBuffer) {
				wg.Add(1)
				go func(buf []string, address string) {
					defer wg.Done()
					shipLinesToUrl(buf, address)
				}(lineBuffer, args[0])
				i = 0
			}
		}
	}
	wg.Wait()
	fmt.Println("Done")
}
