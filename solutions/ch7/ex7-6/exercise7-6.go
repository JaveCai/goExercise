//Exercise 7.6: Add support for Kelvin temperatures to tempflag.

/*
Finish date: 20170716 07:16
*/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.

package main

import (
	"flag"
	"fmt"

	"goExercise/ch7/tempconv"
)

//!+
//var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

var temp = tempconv.KelvinFlag("temp", 293.15, "the temperature in Kelvin")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

//!-
