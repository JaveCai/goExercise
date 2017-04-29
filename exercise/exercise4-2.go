package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		if os.Args[1] == "sha512" {
			s := sha512.Sum512([]byte(os.Args[1]))
			fmt.Printf("%x\n", s)
		} else if os.Args[1] == "sha384" {
			s := sha512.Sum384([]byte(os.Args[1]))
			fmt.Printf("%x\n", s)
		} else if os.Args[1] == "sha256" {
			s := sha256.Sum256([]byte(os.Args[1]))
			fmt.Printf("%x\n", s)
		} else {
			fmt.Printf("usage: ./exec4-2 sha256/384/512\n")
		}
	} else {
		fmt.Printf("usage: ./exec4-2 sha256/384/512\n")
	}

}
