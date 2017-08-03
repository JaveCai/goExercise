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
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/*define a private protocol to send end of list(EOL) to the client*/
var EOL string = "JPROTOCOL:EOL"
var slash string

func main() {

	if runtime.GOOS == "windows" {
		slash = "\\"
	} else { //linux or unix
		slash = "/"
	}

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
	PrevPath := CurPath
	//fmt.Fprintf(conn, "SERVER Serving on %s \r\n", runtime.GOOS)
	in := bufio.NewScanner(conn)
	for in.Scan() {
		fmt.Println(in.Text())

		//strip the command string
		strs := strings.Split(in.Text(), " ")
		if len(strs) == 1 && strs[0] == "close" {
			return
		} else if len(strs) == 1 && strs[0] == "ls" {
			strs = append(strs, ".")
		} else if len(strs) <= 1 {
			fmt.Fprintf(conn, "You need at least one parameter\r\n")
			fmt.Fprintf(conn, EOL)
			fmt.Fprintf(conn, "\r\n")
			continue
		}

		switch strs[0] {
		/* Issue: The second time going into switch, it always matchs default*/
		/* Slove: Figure out the meaning of \r and \n */
		case "ls":
			var files []os.FileInfo
			if strs[1] == "." {
				files, _ = ioutil.ReadDir(CurPath)
			} else if filepath.IsAbs(strs[1]) {
				files, _ = ioutil.ReadDir(strs[1])
			} else {
				files, _ = ioutil.ReadDir(CurPath + slash + strs[1])
			}

			for _, file := range files {
				fmt.Fprintf(conn, file.Name())
				fmt.Fprintf(conn, "\n")

			}
			fmt.Fprintf(conn, "\r")
			fmt.Fprintf(conn, EOL)
			fmt.Fprintf(conn, "\r\n")

		case "cd":
			switch strs[1] {
			case "..":
				err := os.Chdir(getParentDirectory(CurPath))
				if err != nil {
					fmt.Println("err: Invalid path")
				} else {
					PrevPath = CurPath
					CurPath = getParentDirectory(CurPath)
				}
			case "-":
				os.Chdir(PrevPath)
				CurPath = PrevPath

			default: // Generally
				var err error
				if filepath.IsAbs(strs[1]) {
					err = os.Chdir(CurPath)
				} else {
					err = os.Chdir(CurPath + slash + strs[1])
				}

				if err != nil {
					fmt.Println("err: Invalid path")
				} else {
					PrevPath = CurPath
					CurPath = CurPath + slash + strs[1]
				}

			}
			fmt.Fprintf(conn, CurPath+"$ "+"\r\n")

		case "get":
			file, err := os.Open(strs[1]) // For read access.
			if err != nil {
				fmt.Println("err: Open() fail")
				continue
			}
			fileInfo, err := file.Stat()
			if err != nil {
				fmt.Println("err: File.Stat() fail")
				continue
			}
			fmt.Fprintf(conn, "%d\r\n", fileInfo.Size())
			io.CopyN(conn, file, fileInfo.Size())
			file.Close()
			fmt.Fprintf(conn, CurPath+"$ "+"\r\n")

		default:
			fmt.Println("default")
		}

	}

}

func substr(s string, pos1, pos2 int) string {
	runes := []rune(s)
	if pos2-pos1 > 0 && pos2 < len(s) {
		return string(runes[pos1:pos2])
	} else if pos2-pos1 < 0 && pos1 < len(s) {
		return string(runes[pos2:pos1])
	} else {
		return s
	}
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, slash))
}
