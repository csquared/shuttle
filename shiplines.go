package main

import (
	"fmt"
	"net/http"
	"strings"
)

func shipLinesToUrl(lines []string, url string) error {
	body := strings.Trim(strings.Join(lines, ""), "\n")
	reader := strings.NewReader(body)
	resp, err := http.Post(url, "text/plain", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	return nil
}
