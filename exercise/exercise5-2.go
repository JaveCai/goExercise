// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	link, m := visit(nil, nil, doc)
	for key, val := range m {
		fmt.Println("%s: %d", key, val)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, m map[string]int, n *html.Node) ([]string, map[string]int) {

	if n.Type == html.ElementNode {
		m[n.Data]++
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links, m = visit(links, m, c)
	}

	return links, m
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
