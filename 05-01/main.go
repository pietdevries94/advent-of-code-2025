package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(inputFile)

	freshRanges := []Range{}
	freshIDs := 0

	checkIDs := false

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		line := inputScanner.Text()

		if line == "" {
			checkIDs = true
			continue
		}

		if !checkIDs {
			freshRanges = append(freshRanges, RangeStringToRange(line))
			continue
		}

		id, _ := strconv.Atoi(line)
		if IDInAnyRange(id, freshRanges) {
			freshIDs++
		}
	}

	// print the number of zeroes
	fmt.Println(freshIDs)
}

func RangeStringToRange(rangeStr string) Range {
	parts := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return Range{
		start,
		end,
	}
}

func IDInAnyRange(id int, ranges []Range) bool {
	for _, r := range ranges {
		if r.start <= id && r.end >= id {
			return true
		}
	}
	return false
}
