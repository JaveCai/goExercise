/*
Exercise 5.8: Modify forEachNode so that the pre and post functions return a boolean resu lt
indicating whether to continue the traversal. Use it to write a function ElementByID with theSECTION 5.6. ANONYMOUS FUNCTIONS 135
following signature that ﬁnds the ﬁrst HTML element with the speciﬁed id attribute. The
function should stop the traversal as soon as a match is found
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

var key string = "textarea"

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

	ElementByID(doc, key)
	//!+call
	//bGetResult = false
	//forEachNode(doc, startElement, endElement, key)
	//!-call

	return nil
}

var bGetResult bool

func ElementByID(doc *html.Node, id string) *html.Node {
	bGetResult = false
	forEachNode(doc, startElement, endElement, id)
	return doc
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).

func forEachNode(n *html.Node, pre, post func(n *html.Node, key string) bool, key string) {

	if pre != nil && bGetResult == false {
		bGetResult = pre(n, key)
		if bGetResult == true {
			fmt.Printf("pre-----%d\n", bGetResult)
			return
		}
	}

	for c := n.FirstChild; c != nil && bGetResult == false; c = c.NextSibling {
		forEachNode(c, pre, post, key)
	}

	if post != nil && bGetResult == false {
		bGetResult = post(n, key)
		if bGetResult == true {
			fmt.Printf("post-----%d\n", bGetResult)
			return
		}
	}

}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, key string) (ret bool) {
	if n.Type == html.ElementNode {

		if n.FirstChild != nil {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		} else {
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
		}

		if n.Data == key {
			ret = true
			fmt.Printf("-------------------------->\n")
		}

	}
	return ret
}

func endElement(n *html.Node, key string) (ret bool) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)

	}
	if n.Type == html.ElementNode && n.Data == key {
		ret = true
		fmt.Printf("-------------------------->\n")
	}
	return ret
}
