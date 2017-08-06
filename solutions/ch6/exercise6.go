

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

	y.Add(11)
	y.Add(22)
	fmt.Println(y.Len())
	fmt.Println(y.String()) // "{11 22}"

	x.UnionWith(&y)
	fmt.Println(x.Len())
	fmt.Println(x.String()) // "{1 9 11 22 144}"

	x.Remove(9)
	fmt.Println(x.Len())
	fmt.Println(x.String()) // "{1 11 22 144}"

	//x.Clear()
	fmt.Println(x.Len())
	fmt.Println(x.String()) // "{}"

	z := y.Copy()
	z.Add(77)
	fmt.Println(z.Len())
	fmt.Println(z.String()) // "{11 22, 77}"

	fmt.Println(y.Len())
	fmt.Println(y.String()) // "{11 22}"

	fmt.Println(x.Has(9), x.Has(123)) // "false false"

	y.Clear()
	y.AddAll(1, 11, 22, 144, 0)
	fmt.Printf("Elems:")
	s := y.Elems()
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println("")
	//y.Clear()
	fmt.Println(y.Len())
	fmt.Println(y.String())
	//!-main
	fmt.Println("diff:")
	fmt.Println(x.DifferenceWith(&y))

	fmt.Println("same:")
	fmt.Println(x.TheSameWith(&y))

	fmt.Println("intersect:")
	fmt.Println(x.IntersectWith(z))

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
	words []uint
}

/*TODO: exercise6-5*/
//const uintLen = len(uint)

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

/* exercise6-2 */
func (s *IntSet) AddAll(vars ...int) {
	for _, v := range vars {
		s.Add(v)
	}
}

/* exercise6-4 */
func (s *IntSet) Elems() []int {
	sl := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				sl = append(sl, 64*i+j)
			}
		}
	}
	return sl
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

/* exercise6-3 */

func (s *IntSet) IntersectWith(t *IntSet) bool {
	for i, word := range t.words {
		if i < len(s.words) && word == s.words[i] {
			return true
		}
	}
	return false
}

func (s *IntSet) TheSameWith(t *IntSet) bool {

	if len(s.words) == len(t.words) {
		for i, word := range t.words {
			if word != s.words[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) bool {
	return !(s.TheSameWith(t))
}

func (s *IntSet) SymmetricDifference(t *IntSet) bool {
	return !(s.IntersectWith(t))
}

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

/* exercise6-1 */

func (s *IntSet) Len() int { // return the number of elements
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

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word <= len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil

	/*	for i, word := range s.words {
		if word != 0 {
			s.words[i] = nil
		}
	}*/
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
