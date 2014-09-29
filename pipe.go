package main

import (
	"bufio"
	"io"
	"sync"
)

func pipeToHttp(reader io.Reader, url string, bufSize int) {
	buffer := bufio.NewReader(reader)
	lineBuffer := make([]string, bufSize)

	i := 0
	var wg sync.WaitGroup
	for {
		line, err := buffer.ReadString('\n')
		// ship whatever is in the buffer at EOF
		if err == io.EOF {
			_ = shipLinesToUrl(lineBuffer[0:i], url)
			break
		}
		if err != nil {
			panic("Error reading data from buffer")
		}
		// put lines in the buffer; ship them
		// when the buffer is full
		if line != "" {
			lineBuffer[i] = line
			i++
			if i == cap(lineBuffer) {
				wg.Add(1)
				go func(buf []string, _url string) {
					defer wg.Done()
					shipLinesToUrl(buf, _url)
				}(lineBuffer, url)
				i = 0
			}
		}
	}
	wg.Wait()
}
