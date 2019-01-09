package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, Image")
	matrix := createGrid(3, 3)
	fmt.Println(matrix)
	iterations := 1

	for i := 0; i < iterations; i++ {
		matrix = updateGameState(matrix)
		printMatrix(matrix)

	}

}

func printMatrix(matrix [][]int) {
	for y := 0; y < len(matrix); y++ {
		fmt.Println(matrix[y])
	}
}

func createGrid(xLen int, yLen int) [][]int {
	matrix := make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		matrix[i] = make([]int, xLen)
		for j := 0; j < xLen; j++ {
			matrix[i][j] = rand.Intn(2)
		}
	}
	return matrix
}

func updateGameState(matrix [][]int) [][]int {

	yMax := len(matrix)
	xMax := len(matrix[0])

	nextStateMatrix := createGrid(yMax, xMax)

	for i := 0; i < yMax; i++ {
		for j := 0; j < xMax; j++ {
			nextState := getNextCellState(i, j, matrix)
			nextStateMatrix[i][j] = nextState
		}
	}

	return nextStateMatrix
}

func getNextCellState(y int, x int, matrix [][]int) int {
	up := y - 1
	down := y + 1
	left := x - 1
	right := x + 1
	yMax := len(matrix) - 1
	xMax := len(matrix[0]) - 1

	if up < 0 {
		up = 0
	}
	if left < 0 {
		left = 0
	}
	if down >= yMax {
		down = yMax
	}
	if right >= xMax {
		right = xMax
	}

	cellIsLive := matrix[y][x]
	numNeighbors := 0

	fmt.Println("cell analysis:")
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println("up")
	fmt.Println(up)
	fmt.Println("down")
	fmt.Println(down)
	fmt.Println("left")
	fmt.Println(left)
	fmt.Println("right")
	fmt.Println(right)
	for i := up; i <= down; i++ {
		for j := left; j <= right; j++ {
			if !(i == y && j == x) {
				fmt.Println()
				fmt.Println("values")
				fmt.Println(matrix[i][j])
				numNeighbors = numNeighbors + matrix[i][j]
			}
		}
	}
	fmt.Println(cellIsLive)
	fmt.Println(numNeighbors)
	if cellIsLive == 1 {

		//1 and 3
		if numNeighbors < 2 || numNeighbors > 3 {
			return 0
		} else {
			return 1
		}
	} else {
		// Rule 4
		if numNeighbors == 3 {
			return 1
		} else {
			return 0
		}
	}
}
