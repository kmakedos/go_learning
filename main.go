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
	var bits int
	for i := 1; i<8; i++  {
		bits += int(pc[byte(x >> i)])
	}
	return bits
}


func main(){
	flag.Parse()
	for i := range flag.Args() {
		number,_ := strconv.ParseUint(flag.Args()[i],10, 64)
		fmt.Println(PopCount(number))
	}

}