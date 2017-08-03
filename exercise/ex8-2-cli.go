package main

import (
	"bufio"
	"fmt"
	"net"
	"io"
	"os"
	"strings"
)

/*define a private protocol to send EOF to the client*/
var EOF string = "JPROTOCOL:EOF"

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
		if  len(strs) == 1 && strs[0]=="close" {
			return
		}
		switch strs[0] {
		case "ls":
			resp := bufio.NewScanner(conn)
			for resp.Scan() {
				if strings.Contains(resp.Text(), EOF){
					//fmt.Println("EOF")
					break
				}else{
					fmt.Println(resp.Text())
				}
				
			}
		
		case "cd":
			//io.Copy(os.Stdout,conn)
			resp := bufio.NewScanner(conn)
			for resp.Scan() {
				fmt.Println(resp.Text())
				break
			}
	
		case "get":
			file, err:= os.Create(strs[1])
			if err != nil {
				fmt.Println("err: Open fail")
				return
			}
			io.Copy(file, conn)
			file.Close()
			
		}
			
	}

}
