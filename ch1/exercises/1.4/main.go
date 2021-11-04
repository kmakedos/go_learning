package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	f := bufio.NewScanner(os.Stdin)
	for f.Scan(){
		fmt.Println(f.Text())
	}
}
