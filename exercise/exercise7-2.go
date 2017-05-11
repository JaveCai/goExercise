// Copyright © 2017 Jiawei.CAI
// Finish date 201705010

// See page 174. exercise 7.2

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type WordCounter int

func (c WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		c += 1
	}
	return (int)(c), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	nw := w
	c, _ := nw.Write([]byte("1 2"))
	cc := int64(c)
	return nw, &cc
}

func main() {
	var wc WordCounter
	_, pc := CountingWriter(wc) // 2 words
	fmt.Println(*pc)

}
