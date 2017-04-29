package main

import (
	"crypto/sha256"
	"fmt"
	_ "os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	b1 := [32]byte{7, 4}
	b2 := [32]byte{8, 4}
	fmt.Println(bitcomp256(b1, b2))
	s1 := sha256.Sum256([]byte{3, 34})
	s2 := sha256.Sum256([]byte{13, 34})
	fmt.Printf("%x\n%x\n", s1, s2)
	fmt.Println(bitcomp256(s1, s2))
}

func bitcomp256(a, b [32]byte) int {
	var count int = 0
	for i := range a {
		if a[i] != b[i] {
			count += popCount(((uint8)(a[i])) ^ ((uint8)(b[i])))

		}
	}
	return count
}

func popCount(x uint8) int {
	return int(pc[x])
}
