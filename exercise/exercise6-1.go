// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
//package intset

package main

import (
	"bytes"
	"fmt"
)

//func Example_one() {
func main() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.Len())
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.Len())
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.Len())
	fmt.Println(x.String()) // "{1 9 42 144}"

	x.Remove(9)
	fmt.Println(x.Len())
	fmt.Println(x.String()) // "{1 42 144}"

	x.Clear()
	fmt.Println(x.Len())
	fmt.Println(x.String()) // "{}"

	z := y.Copy()
	z.Add(77)
	fmt.Println(z.Len())
	fmt.Println(z.String()) // "{9 42, 77}"

	fmt.Println(y.Len())
	fmt.Println(y.String()) // "{9 42}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word <= len(s.words) {
		s.words[word] &= ^(1 << bit)
	}

}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// return the number of elements
func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i, word := range s.words {
		if word != 0 {
			s.words[i] = 0
		}
	}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var cp IntSet
	cp.words = make([]uint64, len(s.words))
	for i, word := range s.words {
		if word != 0 {
			cp.words[i] = word
		}
	}
	return &cp

}
