package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	_ "path/filepath"
	"strings"
)

const OMDbURL = "http://www.omdbapi.com/"

func main() {
	result, err := SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s\nPoster: %s\n", result.Title, result.Poster)
	saveImage(result)

}

func saveImage(result *Item) {
	resp, err := http.Get(result.Poster)
	if err != nil {
		fmt.Println("HTTP ERROR: %s", err)
		return
	}
	//filename := filepath.Base(poster)
	dst, err := os.Create(result.Title)
	if err != nil {
		fmt.Println("CREAT FILE ERROR: %s", err)
		return
	}

	io.Copy(dst, resp.Body)
}

type Item struct {
	Title  string
	Poster string
}

func SearchIssue(title []string) (*Item, error) {
	q := url.QueryEscape(strings.Join(title, " "))
	resp, err := http.Get(OMDbURL + "?t=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Item
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
