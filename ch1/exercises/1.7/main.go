package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "http://www.in.gr"
	resp, err := http.Get(url)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}

//!-
