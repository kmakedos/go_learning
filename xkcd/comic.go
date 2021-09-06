package xkcd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const XKCD_URL = "https://xkcd.com/"

type ComicStrip struct {
//	Month string
	Num int
//	Link string
//	Year string
//	News string
//	SafeTitle string `json:"safe_title"`
//	Transcript string
//	Alternate string `json: "alt"`
//	Img string
	Title string
//	Day string
}


func Get(num int) (*ComicStrip, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%d/info.0.json", XKCD_URL ,  num ))
	if err != nil {
		log.Printf("Error getting comic strip %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Erroneous Status Code")
	}
	var strip ComicStrip
	if err := json.NewDecoder(resp.Body).Decode(&strip); err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("Could not decode %s",err)
	}
	resp.Body.Close()
	return &strip,nil
}
