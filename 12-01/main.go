package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type shape []*shapeCoordinates

type shapeCoordinates struct {
	x, y int
}

type Area struct {
	ShapeCounts   []int
	Width, Height int
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)

	validAreas := 0
	for inputScanner.Scan() {
		line := inputScanner.Text()

		a, ok := lineToArea(line)
		if !ok {
			continue
		}
		if simpleCheck(a) {
			validAreas++
		}
	}

	fmt.Println(validAreas)
}

func simpleCheck(a *Area) bool {
	// each package can always fit in a 3x3 square
	space := int(a.Height/3) * int(a.Width/3)

	for _, count := range a.ShapeCounts {
		space -= count
		if space < 0 {
			return false
		}
	}
	return true
}

func lineToArea(line string) (*Area, bool) {
	parts := strings.Split(line, ": ")
	if len(parts) == 1 {
		return nil, false
	}

	sizeStrings := strings.Split(parts[0], "x")
	width, _ := strconv.Atoi(sizeStrings[0])
	height, _ := strconv.Atoi(sizeStrings[1])

	shapeCountStrings := strings.Split(parts[1], " ")
	shapeCounts := make([]int, len(shapeCountStrings))
	for i, str := range shapeCountStrings {
		shapeCounts[i], _ = strconv.Atoi(str)
	}

	return &Area{
		ShapeCounts: shapeCounts,
		Width:       width,
		Height:      height,
	}, true
}
