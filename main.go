package main

import (
	"bufio"
	"go_learning/xkcd"
	"log"
	"os"
)



func main(){
	if os.Args[1] == "index"{

		f,err := os.Create("index.txt")
		if err != nil {
			log.Fatalln("Could not create file")
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		for i := 1; i<1000; i++ {
			strip, err := xkcd.Get(i)
			if err != nil {
				log.Printf("Could not download this: %d with error: %s", i, err)
			} else {
				_,err = w.WriteString(strip.Title + "\n")
				if err != nil {
					log.Fatalln("Could not entry write to file")
				}
				if i % 10 == 0  {
					w.Flush()
				}
			}
		}
	}
}