package main

import (
	"fmt"
	"net/http"
	"strings"
)

func shipLinesToUrl(lines []string, url string) error {
	body := strings.NewReader(strings.Join(lines, ""))
	resp, err := http.Post(url, "text/plain", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	return nil
}
