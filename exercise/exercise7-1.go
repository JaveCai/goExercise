// Copyright © 2017 Jiawei.CAI
// Finish date 20170509

// See page 173. exercise 7.1
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//!+bytecounter
type WordCounter int

type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		*c += 1
	}
	return (int)(*c), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	for input.Scan() {
		*c += 1
	}
	return (int)(*c), nil
}

//!-bytecounter

func main() {
	//!+main
	var c WordCounter
	c.Write([]byte("Hello Jave -- 20170509 02:20 can't for asleep"))
	fmt.Println(c)

	var lc LineCounter
	lc.Write([]byte(`1
					 2
					 3
					 4`))
	fmt.Println(lc)

}
