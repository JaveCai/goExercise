// Copyright © 2017 Jiawei.CAI
// Finish date 20170606

// See page 174. exercise 7.2

/*
Exercise 7.2: Write a function CountingWriter with the signature below that, given an
io.Writer,returns a new Writer that wraps the original, and a pointer to an int64 variable
that at any moment contains the number of bytes written to the new Writer.

func CountingWriter(w io.Writer) (io.Writer, *int64)
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		*c += 1
	}
	return (int)(*c), nil
}

type ByteCounter struct {
	Count  int64
	Writer io.Writer
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	bc.Count += int64(len(p)) // convert int to ByteCounter
	return bc.Writer.Write(p)
	//return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var bc ByteCounter
	bc.Writer = w
	return &bc, &(bc.Count)
}

func main() {
	var wc WordCounter
	pw, pc := CountingWriter(&wc)
	pw.Write([]byte("Hello Jave"))
	fmt.Println(wc)
	fmt.Println(*pc)

}
