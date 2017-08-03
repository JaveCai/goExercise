/*
Exercise 7.3: Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4)
that reveals the sequence of values in the tree.
*/

//17.06.08: still not work

package main

import (
	"bytes"
	"fmt"
	//"goExercise/ch4/treesort"
	"math/rand"
	"sort"
	_ "testing"
)

//func TestSort(t *testing.T) {
func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	var tree *Tree
	tree.Sort(data)

	if tree == nil {
		fmt.Println("tree is nil 444")
	} else {
		fmt.Println("tree is available 444")
	}

	if !sort.IntsAreSorted(data) {
		fmt.Errorf("not sorted: %v", data)
	}
	for _, v := range data {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Println(tree.String())
}

//!+
type Tree struct {
	value       int
	left, right *Tree
}

func (root *Tree) String() string {
	values := make([]int, 0)
	var buf bytes.Buffer
	if root == nil {
		fmt.Println("tree is nil 333")
	} else {
		fmt.Println("tree is available 333")
	}

	appendValues(values[:0], root)
	for _, v := range values {
		fmt.Printf("%d ,", v)
	}
	fmt.Println()
	buf.WriteByte('{')
	for _, v := range values {
		fmt.Fprintf(&buf, "%d", v)
		buf.WriteByte(' ')

	}
	buf.WriteByte('}')
	return buf.String()

}

/*
func (root *Tree) Clear() {
	root = nil
}*/

// Sort sorts values in place.
func (root *Tree) Sort(values []int) {
	//var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	if root == nil {
		fmt.Println("tree is nil 111")
	} else {
		fmt.Println("tree is available 111")
	}
	appendValues(values[:0], root)

	if root == nil {
		fmt.Println("tree is nil 222")
	} else {
		fmt.Println("tree is available 222")
	}

}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(Tree)
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
