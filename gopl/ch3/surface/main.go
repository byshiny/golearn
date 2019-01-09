// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
// Need to pipe this as an html file to see the actual picture ./surface > surface.xml

//Byung's notes.
//skipped exercise #2 because I'd have to do  more research on the math to make these weird shapes?
package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
	z float64
}

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			ptA := point{ax, ay, az}
			ptB := point{bx, by, bz}
			ptC := point{cx, cy, cz}
			ptD := point{dx, dy, dz}
			ptList := []point{ptA, ptB, ptC, ptD}
			ctrPt := findCentroid(&ptList)

			fmt.Println("%g, %g, %g, %g", az, bz, cz, dz)
			if areParametersValid(ax, ay, bx, by, cx, cy, dx, dy) {
				if ctrPt.z < 0 {
					fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='blue'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				} else {
					fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='red'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				}
			}

		}
	}
	fmt.Println("</svg>")
}

func findCentroid(ptListPtr *[]point) point {
	var ctrX float64
	var ctrY float64
	var ctrZ float64
	var ptCtr float64
	ptList := *ptListPtr
	for _, v := range ptList {
		ctrX += v.x
		ctrY += v.y
		ctrZ += v.z
		ptCtr++
	}

	ctrPt := point{ctrX / ptCtr, ctrY / ptCtr, ctrZ / ptCtr}
	return ctrPt
}

func areParametersValid(ax float64, ay float64, bx float64, by float64, cx float64, cy float64, dx float64, dy float64) bool {

	if math.IsNaN(ax) {
		return false
	}
	if math.IsNaN(ay) {
		return false
	}
	if math.IsNaN(bx) {
		return false
	}
	if math.IsNaN(by) {
		return false
	}
	if math.IsNaN(cx) {
		return false
	}
	if math.IsNaN(cy) {
		return false
	}
	if math.IsNaN(dx) {
		return false
	}
	if math.IsNaN(dy) {
		return false
	}
	return true
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
