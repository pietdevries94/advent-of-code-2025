package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inOutMap = map[string][]string{}
var pathCache = map[string]counts{}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)

	for inputScanner.Scan() {
		line := inputScanner.Text()

		input, output := ConvertLine(line)
		inOutMap[input] = output
	}

	result := getNumberOfPathsToOut("svr")
	fmt.Println(result.totalboth)
}

type counts struct {
	total     int
	totalfft  int
	totaldac  int
	totalboth int
}

func getNumberOfPathsToOut(input string) counts {
	if input == "out" {
		return counts{
			total: 1,
		}
	}

	if cachedTotal, ok := pathCache[input]; ok {
		return cachedTotal
	}

	outputs, ok := inOutMap[input]
	if !ok {
		return counts{}
	}
	c := counts{}
	for _, output := range outputs {
		outCounts := getNumberOfPathsToOut(output)
		c.total += outCounts.total
		c.totalfft += outCounts.totalfft
		c.totaldac += outCounts.totaldac
		c.totalboth += outCounts.totalboth
	}
	if input == "fft" {
		c.totalfft = c.total
		if c.totaldac > 0 {
			c.totalboth = c.totaldac
		}
	}
	if input == "dac" {
		c.totaldac = c.total
		if c.totalfft > 0 {
			c.totalboth = c.totalfft
		}
	}

	pathCache[input] = c
	return c
}

func ConvertLine(line string) (string, []string) {
	parts := strings.Split(line, ": ")
	return parts[0], strings.Split(parts[1], " ")
}
