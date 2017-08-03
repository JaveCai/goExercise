// Copyright © 2017
// Jave
// Date 20170520

//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	c := make(map[string]int)
	_, count := visit(nil, c, doc)
	for k, v := range count {
		fmt.Printf("%s:\t%d\n", k, v)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, count map[string]int, n *html.Node) ([]string, map[string]int) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.ElementNode {
		//fmt.Println(n.Data)
		switch n.Data {
		case "div", "span", "p":
			count[n.Data]++
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links, count = visit(links, count, c)
	}
	return links, count
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
