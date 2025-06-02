// cmd/pinger/main.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://golang.org",
	}

	for _, url := range urls {
		start := time.Now()

		resp, err := http.Get(url)

		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("error")
			continue
		}
		fmt.Printf("%T\n", resp.Body)
		resp.Body.Close()
		fmt.Printf("Ping %s - %d [%v]\n", url, resp.StatusCode, elapsed)
	}
}
