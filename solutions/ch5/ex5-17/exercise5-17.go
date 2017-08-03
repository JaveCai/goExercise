/*
Exercise 5.17: Wr ite a var iadic function ElementsByTagName that, given an HTML nodetre e
andzeroormorenames, retur nsall the elements thatmatch one ofthose names. Hereare two
examplecal ls:
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
images := ElementsByTagName(doc, "img")
headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
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

var keys []string = []string{"textarea", "a"}

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {

	fmt.Printf("outline----->1\n")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error while calling http.Get()\n")
		return err
	}
	fmt.Printf("outline----->2\n")
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("outline----->3\n")
	/*nodes := */
	/*nodes :=ElementsByTagName(doc, keys...)

	for _,node:=range nodes{
		fmt.Printf("%s\n",node.)
	}*/

	//!+call
	//bGetResult = false
	//forEachNode(doc, startElement, endElement, key)
	//!-call

	return nil
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	//func ElementByID(doc *html.Node, id string) *html.Node {
	var depth int
	nodes := make([]*html.Node, 0)
	fmt.Printf("ElementsByTagName----->1\n")
	forEachNode(
		doc,
		//pre
		func(n *html.Node, keys ...string) {
			if n.Type == html.ElementNode {

				if n.FirstChild != nil {
					fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
					depth++
				} else {
					fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
				}
				for _, v := range keys {
					if v == n.Data {
						nodes = append(nodes, n)
						fmt.Printf("-------------------------->\n")
						break
					}
				}

			}
		},

		func(n *html.Node, keys ...string) {
			if n.Type == html.ElementNode && n.FirstChild != nil {
				depth--
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)

			}
			if n.Type == html.ElementNode {
				for _, v := range keys {
					if v == n.Data {
						nodes = append(nodes, n)
						fmt.Printf("-------------------------->\n")
						break
					}
				}
			}
		},

		name...)
	fmt.Printf("ElementsByTagName----->2\n")
	return nodes
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).

func forEachNode(n *html.Node, pre, post func(n *html.Node, k ...string), keys ...string) {

	if pre != nil {
		pre(n, keys...)
		//fmt.Printf("forEachNode----->pre\n")

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, keys...)

		//fmt.Printf("forEachNode----->for\n")
	}

	if post != nil {
		post(n, keys...)
		//fmt.Printf("forEachNode----->post\n")

	}

}

//!-forEachNode

//!+startend
