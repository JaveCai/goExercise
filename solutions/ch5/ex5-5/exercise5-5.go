//Exercise 5.5: Implement countWordsAndImages.(See Exercise 4.9 for word-splitting.)

// Copyright © 2017 Jave.CAI

// See page 96.

/* Exercis e 4.9: Write a program wordfreq to report the frequency of each word
in an input text ﬁle.Call input.Split(bufio.ScanWords) before the ﬁrst call to
Scan to break the input into words instead of lines.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

//!+



func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)

		
	seen := make(map[string]int) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	input.Split(ScanLetterWords)
	for input.Scan() {
		line := input.Text()
		seen[line] += 1
		//fmt.Println(line)
	}

	s := make(eslice, len(seen))
	i := 0
	for k, c := range seen {
		s[i].str = k
		s[i].count = c
		i++
	}

	sort.Sort(s)
	for _, v := range s {
		fmt.Printf("%d\t: %s\n", v.count, v.str)
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func countWordsAndImages(n *html.Node) (words, images int){
	var images int
	CountImgs(&images,n)

	return words,images
}

func CountImgs(pcount *int, n *html.Node) *int {
	if n.Type == html.ElementNode {
		if n.Data == "img"{
			for _, a := range n.Attr {
				if a.Key == "src" {
					(*pcount)++
				}

			}
		}
		
		

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pcount = visit(pcount, c)
	}
	return pcount
}


type e struct {
	str   string
	count int
}

type eslice []e

func (p eslice) Len() int           { return len(p) }
func (p eslice) Less(i, j int) bool { return p[i].count > p[j].count }
func (p eslice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func ScanLetterWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading rune which is not letter.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if unicode.IsLetter(r) {
			break
		}
	}
	// Scan until notspace, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !unicode.IsLetter(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
