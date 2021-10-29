package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func fib1(n int) int {
	if n < 2 {
		return n
	} else {
		return fib1(n-1) + fib1(n-2)
	}
}


func fib2(n int) int {
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i<n+1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}


func main(){
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments, please give me a num")
	}
	n,err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	start := time.Now()
	fmt.Println("Fib1:", fib1(n))
	log.Println("Fib1 took:", time.Since(start))
	start = time.Now()
	fmt.Println("Fib2:", fib2(n))
	fmt.Println("Fib2 took:", time.Since(start))
}