package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fib1(n int) int {
	if n <= 1 {
		return 0
	} else {
		return fib1(n-1) + fib1(n-2)
	}
}


func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter num>")
	text,_ := reader.ReadString('\n')
	n,err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(fib1(n))
}