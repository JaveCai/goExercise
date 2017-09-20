/*
Exercise 8.14: Change the chat server’s network protocol so that each client provides its name
on entering. Use that name instead of the network address when prefixing each message with
its sender’s identity.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
	"strings"
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
	ch <- "Please enter your name to login"
	input := bufio.NewScanner(conn)
	input.Scan() 

	who = strings.Join([]string{input.Text(),who},"@")
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
	input = bufio.NewScanner(conn)
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
