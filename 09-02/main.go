package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords = struct {
	x, y int
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)

	allCoords := []Coords{}

	matrixWidth := 0
	matrixHeight := 0

	for inputScanner.Scan() {
		line := inputScanner.Text()

		current := StringToCoords(line)
		allCoords = append(allCoords, current)

		if current.x > matrixWidth {
			matrixWidth = current.x
		}
		if current.y > matrixHeight {
			matrixHeight = current.y
		}
	}

	var prev *Coords

	matrix := BuildMatrix(matrixWidth, matrixHeight)
	for i, current := range allCoords {
		fmt.Printf("Building Matrix: %v/%v\n", i, len(allCoords))
		if i == 0 {
			prev = &allCoords[len(allCoords)-1]
		}

		// draw a vertical line
		if prev.x == current.x {
			start := min(prev.y, current.y)
			end := max(prev.y, current.y)

			x := current.x
			for y := start; y <= end; y++ {
				matrix[x][y] = 1
			}
		}

		// draw a horizontal line
		if prev.y == current.y {
			start := min(prev.x, current.x)
			end := max(prev.x, current.x)

			y := current.y
			for x := start; x <= end; x++ {
				matrix[x][y] = 1
			}
		}

		prev = &current
	}

	allPreviousCords := []Coords{}
	largestSize := 0

	for i, current := range allCoords {
		fmt.Printf("Finding: %v/%v\n", i, len(allCoords))
		for _, c := range allPreviousCords {
			size := CalculateSize(current, c)
			if size > largestSize && SquareInMatrix(matrix, current, c) {
				largestSize = size
			}
		}

		allPreviousCords = append(allPreviousCords, current)
	}

	fmt.Println(largestSize)
}

func StringToCoords(str string) Coords {
	parts := strings.Split(str, ",")
	c := Coords{}
	c.x, _ = strconv.Atoi(parts[0])
	c.y, _ = strconv.Atoi(parts[1])
	return c
}

func CalculateSize(a, b Coords) int {
	x := GetDiff(a.x, b.x)
	y := GetDiff(a.y, b.y)

	return x * y
}

// We add a + 1 because we want to include both coords
func GetDiff(a, b int) int {
	if a > b {
		return a - b + 1
	}
	return b - a + 1
}

func BuildMatrix(matrixWidth, matrixHeight int) [][]int {
	res := make([][]int, matrixWidth+1)
	for i := range matrixWidth + 1 {
		res[i] = make([]int, matrixHeight+1)
	}
	return res
}

func SquareInMatrix(matrix [][]int, a, b Coords) bool {
	// I make the educated guess that if the diff between the x's or y's is less than 3, it won't be the answer
	// This simplifies the implementation
	if GetDiff(a.x, b.x) < 3 || GetDiff(a.y, b.y) < 3 {
		return false
	}

	minX := min(a.x, b.x) + 1
	maxX := max(a.x, b.x) - 1
	minY := min(a.y, b.y) + 1
	maxY := max(a.y, b.y) - 1

	for y := minY; y <= maxY; y++ {
		if matrix[minX][y] == 1 {
			return false
		}
		if matrix[maxX][y] == 1 {
			return false
		}
	}

	for x := minX; x <= maxX; x++ {
		if matrix[x][minY] == 1 {
			return false
		}
		if matrix[x][maxY] == 1 {
			return false
		}
	}
	return true
}
