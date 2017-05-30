/*
Exercise 5.7:
Develop startElement and endElement into a general HTML pretty-printer.
Print comment nodes, text nodes, and the attributes of each element(<a href='...'>). Use
short forms like <img/> instead of <img></img> when an element has no children. Write a
test to ensure that the output can be parsed successfully. (See Chapter 11.)
*/

// base on outline2
// date 2017.05.29

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {

		if n.FirstChild != nil {

			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)

			if n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						fmt.Printf("%*s<%s>\n", depth*2, "", a.Val)
						//links = append(links, a.Val)
					}
				}
			}

			depth++

		} else {
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
		}

	}

}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)

	}
}
