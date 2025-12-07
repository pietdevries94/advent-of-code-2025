package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(inputFile)

	numberOfSpliters := 0
	beamIndexes := map[int]bool{}

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		line := inputScanner.Text()

		maxLineIndex := len(line)

		start := strings.Index(line, "S")
		if start != -1 {
			beamIndexes[start] = true
		}

		newBeamIndexes := map[int]bool{}
		for beamIndex := range beamIndexes {
			if line[beamIndex] != '^' {
				newBeamIndexes[beamIndex] = true
				continue
			}

			numberOfSpliters++

			if beamIndex > 0 {
				newBeamIndexes[beamIndex-1] = true
			}
			if beamIndex < maxLineIndex {
				newBeamIndexes[beamIndex+1] = true
			}

		}
		beamIndexes = newBeamIndexes
	}

	fmt.Println(numberOfSpliters)
}
