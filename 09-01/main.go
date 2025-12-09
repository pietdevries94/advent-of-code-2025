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

	allPreviousCords := []Coords{}
	largestSize := 0

	for inputScanner.Scan() {
		line := inputScanner.Text()

		current := StringToCoords(line)
		for _, c := range allPreviousCords {
			size := CalculateSize(current, c)
			if size > largestSize {
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
