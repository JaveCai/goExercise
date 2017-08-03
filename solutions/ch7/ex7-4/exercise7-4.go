/*
Exercise 7.4: The strings.NewReader function returns a value that satisﬁes the io.Reader
interface (and others) by reading from its argument, a string . Implement a simple version of
NewReader yourself,and use it to make the HTML parser(§5.2) take input from a string .
*/

//date:2017.07.06
//status:coding

package main 

import(
	"fmt"
	"net/http"
	"os"
)

type Reader{
	s string
}

func (r *Reader)Read(b []byte)(n int,err error){
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(b, resp.Body[0:])
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return n, err
}

func NewReader(s string) *Reader{
	return &Reader{s}
}

func Parse(url string){
	r:=NewReader(url)
	var o []byte
	r.Read(o)
	fmt.Println(o)
}