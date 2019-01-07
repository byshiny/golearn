// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("herro")
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http://")
		if !hasPrefix {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Println(err)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		statusReader := strings.NewReader(resp.Status)
		_, err = io.Copy(os.Stdout, statusReader)
		// b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()c
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		//fmt.Printf("%s", b)
	}
}

//!-
