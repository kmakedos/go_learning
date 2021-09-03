package xkcd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const XKCD_URL = "https://xkcd.com/"

type ComicStrip struct {
	Month string
	Num int
	Link string
	Year string
	News string
	SafeTitle string `json:"safe_title"`
	Transcript string
	Alternate string `json: "alt"`
	Img string
	Title string
	Day string
}



func Get(num string){
	resp, err := http.Get(XKCD_URL + num + "/info.0.json")
	if err != nil {
		log.Fatalln("Error getting comic strip")
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Some error occurred")
		resp.Body.Close()
	}
	var strip ComicStrip
	if err := json.NewDecoder(resp.Body).Decode(&strip); err != nil {
		log.Fatalf("Error decoding %s", err)
		resp.Body.Close()
	}
	resp.Body.Close()
	fmt.Println(strip.Title)
}
