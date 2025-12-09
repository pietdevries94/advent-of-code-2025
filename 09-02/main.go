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

	allCoords := []*Coords{}

	for inputScanner.Scan() {
		line := inputScanner.Text()

		current := StringToCoords(line)
		allCoords = append(allCoords, &current)
	}

	allPreviousCords := []*Coords{}
	largestSize := 0

	for _, current := range allCoords {
		for _, c := range allPreviousCords {
			size := CalculateSize(current, c)
			if size > largestSize && SquareInMatrix(allCoords, current, c) {
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

func CalculateSize(a, b *Coords) int {
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

func SquareInMatrix(allCoords []*Coords, a, b *Coords) bool {
	minX := min(a.x, b.x)
	maxX := max(a.x, b.x)
	minY := min(a.y, b.y)
	maxY := max(a.y, b.y)

	var prev *Coords
	for i, c := range allCoords {
		if i == 0 {
			prev = allCoords[len(allCoords)-1]
		} else {
			prev = allCoords[i-1]
		}

		// if the x is the same, we need to check if the a is in the y range of the line
		if prev.x == c.x && c.x > minX && c.x < maxX {
			minYLine := min(prev.y, c.y)
			maxYLine := max(prev.y, c.y)
			if maxY > minYLine && minY < maxYLine {
				return false
			}
		}

		// the same for y
		if prev.y == c.y && c.y > minY && c.y < maxY {
			minXLine := min(prev.x, c.x)
			maxXLine := max(prev.x, c.x)
			if maxX > minXLine && minX < maxXLine {
				return false
			}
		}
	}
	return true
}
