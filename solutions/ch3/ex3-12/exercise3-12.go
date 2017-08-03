/*

Date: 20170326
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var u uint8 = 0x55
	fmt.Printf("Hello, sublime. %b\n", ^u)
	fmt.Println(anagrams(os.Args[1], os.Args[2]))

}
func anagrams(s1, s2 string) bool {
	if contain(s1, s2) && contain(s2, s1) {
		return true
	}
	return false
}

func contain(s1, s2 string) bool {
	//var s []string = s1
	for i := range s1 {
		index := strings.IndexAny(s1[i:], s2)
		if index != 0 {
			return false
		}
	}
	return true

}
