package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	for _,url := range os.Args[1:]{
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		n,err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\nWritten %d bytes\n", n)
		fmt.Printf("Status Code: %s\n", resp.Status)
	}
}