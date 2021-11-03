package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	sep = " "
	s = strings.Join(os.Args, sep)
	fmt.Println(s)
}


