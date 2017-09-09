/*
Exercise 8.8: Using a select statement, add a timeout to the echo server from Section 8.3 so
that it disconnects any client that shouts nothing within 10 seconds.
*/

/*
Solution: add a "watch dog" for each client goroutine
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	maxIdleTime := 10
	// wg+
	feed := make(chan bool)
	go func(c net.Conn, f chan bool) {
		secTicker := time.NewTicker(time.Second)
		wg := 0
		for {
			select {
			case <-f:
				fmt.Println("reset tick count.")
				wg = 0
			case <-secTicker.C:
				fmt.Println("tick...")
				wg += 1
				if wg >= maxIdleTime {
					msg := c.RemoteAddr().String() + ": idle too long, this link was closed."
					fmt.Println(msg)
					fmt.Fprintln(c, msg)
					secTicker.Stop()
					c.Close()
					return
				}

			}
		}
	}(c, feed)

	//wg-
	for input.Scan() {
		feed <- true // feed the watch dog
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
