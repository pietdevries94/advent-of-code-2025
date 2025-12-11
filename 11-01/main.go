package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var InOutMap = map[string][]string{}
var PathCache = map[string]int{}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)

	for inputScanner.Scan() {
		line := inputScanner.Text()

		input, output := ConvertLine(line)
		InOutMap[input] = output
	}

	result := GetNumberOfPathsToOut("you")
	fmt.Println(result)
}

// type Counts

func GetNumberOfPathsToOut(input string) int {
	if input == "out" {
		return 1
	}
	if cachedTotal, ok := PathCache[input]; ok {
		return cachedTotal
	}
	outputs, ok := InOutMap[input]
	if !ok {
		return 0
	}
	total := 0
	for _, output := range outputs {
		total += GetNumberOfPathsToOut(output)
	}
	PathCache[input] = total
	return total
}

func ConvertLine(line string) (string, []string) {
	parts := strings.Split(line, ": ")
	return parts[0], strings.Split(parts[1], " ")
}
