package main

import (
	"flag"
	"fmt"
	"go_learning/temperature"
	"strconv"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var c = flag.Bool("c", false, "Celsius values")
var f = flag.Bool("f", false, "Fahrenheit values")

func main(){
	flag.Parse()
		for i := range flag.Args() {
			number, _ := strconv.ParseFloat(flag.Args()[i], 64)
			if *c {
				fmt.Printf("%2.2f oC\n", temperature.CtoF(temperature.Celsius(number)))
			}
			if *f {
				fmt.Printf("%2.2f oF\n", temperature.FtoC(temperature.Fahrenheit(number)))
			}
		}
}