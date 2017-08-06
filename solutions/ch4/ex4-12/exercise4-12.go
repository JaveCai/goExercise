
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const IssuesURL = "https://xkcd.com/"

type comic struct {
	Title      string
	Transcript string
	Img        string
}

//!+
func main() {

	resp, err := http.Get(IssuesURL + os.Args[1] + "/info.0.json")
	if err != nil {
		fmt.Println("Http Get failed: %s", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println("Response Status error: %s", resp.Status)
		return
	}

	var ret comic
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		resp.Body.Close()
		fmt.Println("Decoder error: %s", err)
		return
	}
	resp.Body.Close()

	//result, err := GetComic(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\nTranscript: %s\nImg: %s\n",
		ret.Title, ret.Transcript, ret.Img)
}

/*
func GetComic(index string)(*comic, error){
	resp, err := http.Get(IssuesURL + index + "/info.0.json")
	if err != nil {
		return nil,err
	}
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("Get comic failed: %s", resp.Status)
	}

	var ret comic
	if err:=json.NewDecoder(resp.Body).Decode(&ret); err!=nil {
		resp.Body.Close()
		return nil,err
	}
	resp.Body.Close()
	return &ret, nil
}*/
