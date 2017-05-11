package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	_ "testing"
	//"gopl.io/ch4/treesort"
)

//func TestSort(t *testing.T) {
func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		fmt.Errorf("not sorted: %v", data)
	}
	for _, v := range data {
		fmt.Printf("%d ", v)
	}
}

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var buf bytes.Buffer
	buf.Write("{")

}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	root.String()
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
