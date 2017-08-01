/*
go ftp server
1.listen and accept and goroutine
2.read the input from terminal
3.deal with command and send the response
*/

package main 

import(
	"fmt"
	"net"
	"strings"
	"bufio"
	"io/ioutil"
	"os"
	_"path/filepath"
)

func main(){

	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept err")// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	fmt.Fprintf(conn, "SERVER SERVING...\r\n\r")
	in:=bufio.NewScanner(conn)
	for in.Scan(){
		fmt.Println(in.Text())

		//strip the command string
		strs:=strings.Split(in.Text()," ") 
		if len(strs)==1 && strings.Contains(strs[0],"close"){
			return
		}else if len(strs)<=1 {
			fmt.Println("You need at least one parameter")
			continue
		}
		//fmt.Println(strs[0])

		switch strs[0]{
			case "ls":
				files,_:=ioutil.ReadDir(strs[1])
				for _,file:=range files{
					fmt.Println(file.Name())
					fmt.Fprintf(conn,file.Name())
					fmt.Fprintf(conn,"\r\n\r")
				}
/*
				fs,_ :=filepath.Glob(strs[1])
				fmt.Println(fs)*/

			case "cd": fmt.Println("You arequest for cd")
				os.Chdir(strs[1])
			case "get": fmt.Println("You arequest for get")

		}
	}

}

