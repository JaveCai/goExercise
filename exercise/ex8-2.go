/*
go ftp server
1.listen and accept and goroutine
2.read the input from terminal
3.deal with command and send the response back to client
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
)

/*define a private protocol to send EOF to the client*/
var EOF string = "JPROTOCOL:EOF"

func main() {

	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept err") // handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	CurPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	PrevPath:=CurPath
	//fmt.Fprintf(conn, "SERVER Sserving from %s \r\n", CurPath)
	in := bufio.NewScanner(conn)
	for in.Scan() {
		fmt.Println(in.Text())

		//strip the command string
		strs := strings.Split(in.Text(), " ")
		if len(strs) == 1 && strs[0]=="close" {
			return
		} else if len(strs) <= 1 {
			fmt.Fprintf(conn, "You need at least one parameter") 
			fmt.Fprintf(conn, "\r\n")
			fmt.Fprintf(conn, EOF)
			fmt.Fprintf(conn, "\r\n")
			continue
		}
		//fmt.Println(strs[0])

		switch strs[0] {
		/* Issue: The second time going into switch, it always matchs default*/
		/* Slove: Figure out the meaning of \r and \n */
		case "ls":
			var files []os.FileInfo
			if strs[1]=="."{
				files, _ = ioutil.ReadDir(CurPath)
			}else if filepath.IsAbs(strs[1]){
				files, _ = ioutil.ReadDir(strs[1])
			}else{
				files, _ = ioutil.ReadDir(CurPath + `\` + strs[1])
			}
			
			for _, file := range files {
				//fmt.Println(file.Name())

				fmt.Fprintf(conn, file.Name())
				fmt.Fprintf(conn, "\n")

			}

			fmt.Fprintf(conn, "\r")
			fmt.Fprintf(conn, "%s$\r\n",CurPath) 
			/*define a private protocol to send EOF to the client*/
			fmt.Fprintf(conn, EOF)
			fmt.Fprintf(conn, "\r\n")


		case "cd":
			fmt.Println("You arequest for cd")
			switch strs[1]{

			case "..":
				err:= os.Chdir(getParentDirectory(CurPath))
				if err != nil{
					fmt.Println("err: Invalid path")
				}else{
					PrevPath = CurPath
					CurPath = getParentDirectory(CurPath)
				}

			case "-":
				os.Chdir(PrevPath)
				CurPath = PrevPath

			default: // Generally 
				var err error
				if filepath.IsAbs(strs[1]){
					err= os.Chdir(CurPath)
				}else{
					err= os.Chdir(CurPath + `\` + strs[1])
				}
				
				if err != nil{
					fmt.Println("err: Invalid path")
				}else{
					PrevPath=CurPath
					CurPath=CurPath + `\` + strs[1]
				}
					
			}
			fmt.Fprintf(conn, CurPath+"$") 
			fmt.Fprintf(conn, "\r\n") 
		case "get":
			fmt.Println("You arequest for get")

		default:
			fmt.Println("default")
		}
	}

}

func substr(s string, pos1, pos2 int) string {
	runes := []rune(s)
	if pos2-pos1>0 && pos2<len(s){
		return string(runes[pos1:pos2])
	} else if pos2-pos1<0 && pos1<len(s){
		return string(runes[pos2:pos1])
	}else{
		return s
	}
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "\\"))
}
