package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type BeamIndexMap = map[int]bool
type BeamIndexStr = string

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(inputFile)

	beamIndexes := map[BeamIndexStr]int{}

	lineNumber := 0

	for inputScanner.Scan() {
		line := inputScanner.Text()

		beamIndexes = ProcessLine(line, beamIndexes)

		lineNumber++
		fmt.Printf("parsed line: %v\n", lineNumber)
	}

	total := 0
	for _, count := range beamIndexes {
		total += count
	}

	fmt.Println(total)
}

func ProcessLine(line string, beamIndexes map[BeamIndexStr]int) map[BeamIndexStr]int {
	start := strings.Index(line, "S")
	if start != -1 {
		newBeamIndexes := BeamIndexMap{
			start: true,
		}
		beamIndexes = map[BeamIndexStr]int{
			serializeBeamIndexes(newBeamIndexes): 1,
		}
		return beamIndexes
	}

	newBeamIndexes := map[BeamIndexStr]int{}
	for bi, numberOfTimelines := range beamIndexes {
		bim := deserializeBeamIndexes(bi)
		res := ProcessTimeline(line, bim, numberOfTimelines)
		for str, count := range res {
			newBeamIndexes[str] = newBeamIndexes[str] + count
		}
	}

	return newBeamIndexes
}

func ProcessTimeline(line string, beamIndexes BeamIndexMap, count int) map[BeamIndexStr]int {
	maxLineIndex := len(line)

	allNewBeamIndexes := map[BeamIndexStr]int{}

	for beamIndex := range beamIndexes {
		if line[beamIndex] != '^' {
			addToAll(allNewBeamIndexes, beamIndex, count)
			continue
		}

		if beamIndex > 0 {
			addToAll(allNewBeamIndexes, beamIndex-1, count)
		}
		if beamIndex < maxLineIndex {
			addToAll(allNewBeamIndexes, beamIndex+1, count)
		}
	}

	return allNewBeamIndexes
}

func addToAll(target map[BeamIndexStr]int, index int, count int) {
	m := BeamIndexMap{
		index: true,
	}
	str := serializeBeamIndexes(m)
	target[str] = target[str] + count
}

func serializeBeamIndexes(bi BeamIndexMap) BeamIndexStr {
	res, _ := json.Marshal(bi)
	return string(res)
}

func deserializeBeamIndexes(str BeamIndexStr) BeamIndexMap {
	var res BeamIndexMap
	json.Unmarshal([]byte(str), &res)
	return res
}
