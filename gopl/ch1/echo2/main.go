// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	s, sep := "", ""

	for idx, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(idx, arg)
	}
	duration := time.Since(start).Seconds()
	fmt.Println("loop option")
	fmt.Println(s)
	fmt.Println(duration)
	fmt.Println("built in library option")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	duration = time.Since(start).Seconds()
	fmt.Println(duration)
	fmt.Println("OMG. A Magnitude order of difference!")
}

//!-
