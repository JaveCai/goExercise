package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Dial err") // handle error
	}
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r")
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Fprintf(conn, in.Text())
		fmt.Fprintf(conn, "\r\n\r") //send msg
		strs := strings.Split(in.Text(), " ")
		if len(strs) == 1 && strings.Contains(strs[0], "close") {
			return
		}
		switch strs[0] {
		case "ls":
			resp := bufio.NewScanner(conn)
			for resp.Scan() {
				fmt.Println(resp.Text())
			}
		}
	}

}
