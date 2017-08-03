/*
go ftp client for ex8.2
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

/*define a private protocol to send EOL*/
var EOL string = "JPROTOCOL:EOL"

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Dial err") // handle error
	}
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Fprintf(conn, in.Text())
		fmt.Fprintf(conn, "\r\n") //send msg
		strs := strings.Split(in.Text(), " ")
		if len(strs) == 1 && strs[0] == "close" {
			return
		}
		switch strs[0] {
		case "ls":
			resp := bufio.NewScanner(conn)
			for resp.Scan() {
				if strings.Contains(resp.Text(), EOL) {
					break
				} else {
					fmt.Println(resp.Text())
				}

			}

		case "cd":
			resp := bufio.NewScanner(conn)
			resp.Scan()
			fmt.Printf(resp.Text())

		case "get":
			var size int
			resp := bufio.NewScanner(conn)
			resp.Scan()
			size, _ = strconv.Atoi(resp.Text())
			//fmt.Printf("file size: %d\n", size)

			file, err := os.Create(strs[1])
			if err != nil {
				fmt.Println("err: Create file fail")
				return
			}
			io.CopyN(file, conn, int64(size))
			file.Close()
			fmt.Printf("Done!\n")

			resp.Scan()
			fmt.Printf(resp.Text())

		}

	}

}
