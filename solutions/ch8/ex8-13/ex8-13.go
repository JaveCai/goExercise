/*
Exercise 8.13: Make the chatser ver disconne ctidleclients, such as those that have sent no
messages in the last ﬁve minutes(30 sec). Hint: cal ling conn.Close() in another goroutine unblocks
active Read calls such as the one done by input.Scan().
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client struct {
	Out  chan<- string // an outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

var maxIdleTime int = 30 // seconds

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.Out <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			for c := range clients {
				cli.Out <- c.Name
			}
			cli.Out <- "are here!"

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	cli := client{ch, who}
	messages <- who + " has arrived"
	entering <- cli

	idleTime := 0
	go func() {
		secTicker := time.NewTicker(time.Second)
		for {
			select {
			case <-secTicker.C:
				idleTime++
				if idleTime == maxIdleTime {
					conn.Close()
					secTicker.Stop()
					break
				}
			}
		}
	}()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		idleTime = 0
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
