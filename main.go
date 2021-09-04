package main

import (
	"bufio"
	"go_learning/xkcd"
	"log"
	"os"
)


func check(e error){
	if e != nil {
		log.Println(e)
	}

}
func check_fatal(e error){
	if e != nil {
		panic(e)
	}
}

func main(){
	if os.Args[1] == "index"{

		f,err := os.Create("index.txt")
		check_fatal(err)
		defer f.Close()
		w := bufio.NewWriter(f)
		for i := 1; i<1000; i++ {
			strip, err := xkcd.Get(i)
			check(err)
			_,err = w.WriteString(strip.Title + "\n")
			check(err)
			if i % 10 == 0  {
				w.Flush()
			}
		}
	}
}