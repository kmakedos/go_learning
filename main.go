package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func fib1(n int) int {
	if n < 2 {
		return n
	} else {
		return n + fib1(n-1)
	}
}


func main(){
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments, please give me a num")
	}
	n,err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(fib1(n))
}