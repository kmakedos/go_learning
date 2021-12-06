package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := []string{"www.in.gr", "http://www.google.gr"}
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, err2 := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err2 != nil {
				return
			}
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
