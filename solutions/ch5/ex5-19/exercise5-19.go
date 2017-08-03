/*
Exercise 5.19: Use panic and recover to write a function that contains no return statement
yet returns a non-zero value.

date:17.06.07
*/
package main

import "fmt"

type bail struct{}

func main() {

	fmt.Println(test())
}

func test() (ret int) {
	defer func() {
		switch p := recover(); p {
		case bail{}:
			ret = 5
		default:
			panic(p)
		}
	}()
	panic(bail{})

}
