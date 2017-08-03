// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strs []string) []string {
	i := 0
	prev := ""
	for _, s := range strs {
		if i == 0 {
			strs[i] = s
			prev = s
			i++
		} else if s != prev {
			strs[i] = s
			prev = s
			i++

		}
	}
	return strs[:i]
}

//!-nonempty

func main() {
	//!+main
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var strs []string
		for _, s := range strings.Fields(input.Text()) {
			strs = append(strs, s)
		}
		//reverse(ints)
		fmt.Printf("%q\n", nonempty(strs))
	}
	//!-main
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//!-alt
