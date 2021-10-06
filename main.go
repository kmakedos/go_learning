package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"go_learning/xkcd"
	"log"
	"os"
)


func updateMap(numbers []int) map[int]int{
	m := make(map[int]int)
	for _,v := range numbers{
		m[v]++
	}
	return m
}

func main(){
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}
	wordMap := make(map[string][]int)
	if os.Args[1] == "index"{
		f,err := os.Create("index.txt")
		if err != nil {
			log.Fatalln("Could not create file")
		}
		defer f.Close()
		for i := 1; i<2000; i++ {
			strip, err := xkcd.Get(i)
			if err != nil {
				log.Printf("Could not download this: %d with error: %s", i, err)
			} else {
				for _, word := range xkcd.Tokenize(strip) {
					wordMap[word] = append(wordMap[word], strip.Num)
				}
			}
		}

		w := bufio.NewWriter(f)
		e := gob.NewEncoder(w)
		err = e.Encode(wordMap)
		if err != nil {
			log.Fatalln(err)
		}
		w.Flush()
	} else {
		f, err := os.Open("index.txt")
		if err != nil {
			log.Fatalln(err)
		}
		d := gob.NewDecoder(f)
		err = d.Decode(&wordMap)
		if err != nil {
			log.Fatalln(err)
		}
		var ids []int
		for _, word := range os.Args[1:] {
			if m, ok := wordMap[word]; ok {
				ids = append(ids, m...)
			}
		}
		for _,id := range ids {
			fmt.Printf("%s%d\n", xkcd.XKCD_URL, id)
		}
	}

}