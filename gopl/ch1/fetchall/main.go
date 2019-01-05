// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	topWebsites := readCsvFile("top-1m.csv")

	start := time.Now()
	ch := make(chan string)
	// os.Args[1:]
	for _, websiteRecord := range topWebsites {
		url := websiteRecord[1]
		go fetch(url, ch) // start a goroutine
	}
	for range topWebsites {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func readCsvFile(filePath string) [][2]string {
	// Load a csv file.
	f, _ := os.Open(filePath)
	recordList := make([][2]string, 100)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		//fmt.Println(record)
		// for value := range record {
		// 	fmt.Printf("  %v\n", record[value])
		// }
		websiteLink := "https://www." + record[1]
		recordList = append(recordList, [2]string{record[0], websiteLink})
	}

	return recordList
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	//watch out for https vs http - need to take care of http later.
	//this is because when you write to the file it doesn't like slashes
	httpStrippedFilename := strings.Replace(url, "https://", "", 1)
	urlname := httpStrippedFilename + ".html"
	directoryName := "./websites"
	fullPath := path.Join(directoryName, urlname)
	fmt.Println(fullPath)
	to, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	nbytes, err := io.Copy(to, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
