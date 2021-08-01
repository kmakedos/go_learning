package main

import (
	"bufio"
	"fmt"
	"os"
)

type info struct {
	counts map[string]int
	filename string
}


func main(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d %s\n", n, line)
			}
		}
	} else {
		for _,arg := range files {
			f,err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr,"dup2: %v\n", err)
			}
			countLines(f, counts)
			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%s %d %s\n", f.Name(), n, line)
				}
			}
			f.Close()
		}

	}

}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}
