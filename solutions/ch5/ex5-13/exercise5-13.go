/*
Exercise 5.13: Modify crawl to makelocal copies ofthe pages it ﬁnds,creating directories as
necessary.Don’t make copies of pages that come from a different domain. For example, if the
original page comes from golang.org,save all ﬁles from there, but exclude ones from
vimeo.com.
*/
// See page 139.

//date: 2017.06.03

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"io"
	"net/http"
	"path"

	"goExercise/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	url := worklist[0]
	fmt.Println(url)
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] && strings.Contains(item, url) {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)

	local, n, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)

	} else {
		//fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
		fmt.Printf("%s => %s (%d bytes).\n", url, local, n)
	}

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	} else if strings.Contains(local, ".html") == false {
		local = local + ".html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
