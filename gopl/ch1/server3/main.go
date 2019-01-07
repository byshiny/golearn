// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	lisaHandler := func(w http.ResponseWriter, r *http.Request) {
		cycleInput := r.URL.Query().Get("cycles")
		cycles := 5
		if cycleInput != "" {
			cyclesConverted, err := strconv.Atoi(cycleInput)
			if err != nil {
				log.Fatal("Cycles need to be an integer")
			}
			cycles = cyclesConverted
		}

		lissajous(w, cycles)
	}
	http.HandleFunc("/lisa", lisaHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//Don't understand enough about package management - book doesn't explain this ugh in Ch. 1
const (
	whiteIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	redIndex   = 2
	blueIndex  = 3
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}}

func lissajous(out io.Writer, cycleInput int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	cycles := float64(cycleInput) // number of complete x oscillator revolutions
	freq := rand.Float64() * 3.0  // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	//colorIndex := redIndex
	looper := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if looper == 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					redIndex)
			} else if looper == 1 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					greenIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					blueIndex)
			}
		}
		looper = (looper + 1) % 3

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-handler
