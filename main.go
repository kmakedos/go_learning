package main

import (
	"flag"
	"fmt"
	"strconv"
)

var pc [256]byte
func init(){
	for i:= range pc{
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}


func main(){
	flag.Parse()
	for i := range flag.Args() {
		number,_ := strconv.ParseUint(flag.Args()[i],10, 64)
		fmt.Println(PopCount(number))
	}

}