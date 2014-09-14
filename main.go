package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	batchSize := 5
	lineBuffer := make([]string, batchSize)
	i := 0
	args := os.Args[1:]

	for {
		line, err := stdin.ReadString('\n')
		if err == io.EOF {
			_ = shipLinesToUrl(lineBuffer, args[0])
			break
		}
		if err != nil {
			panic("Error reading data from stdin")
		}
		if line != "" {
			lineBuffer[i] = line
			i++
			if i == batchSize {
				go shipLinesToUrl(lineBuffer, args[0])
				i = 0
			}
		}
	}
}
