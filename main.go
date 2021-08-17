package main

import (
	"fmt"
)

func arcopy(x [3]int) [3]int {
	x[0] = -1
	return x
}

func achange(x [3]int) {
	x[0] = -2
}

func slicechange(sx []int){
	sx[0] = -3
}
func main(){
	x := [3]int{1,2,3}
	for k := range x {
		fmt.Println( x[k])
	}
	arcopy(x)
	for k:=range x{
		fmt.Println(x[k])
	}
	achange(x)
	for k:=range x{
		fmt.Println(x[k])
	}
	slicechange(x[0:2])
	for k:= range x{
		fmt.Println(x[k])
	}

}