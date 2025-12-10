package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type System struct {
	indicatorTarget []bool
	buttons         [][]int
	joltage         []int
}

type Coords = struct {
	x, y int
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)

	total := 0

	for inputScanner.Scan() {
		line := inputScanner.Text()

		total += ParseLine(line)
	}

	fmt.Println(total)
}

func ParseLine(line string) int {
	system := LineToSystem(line)
	return BruteForceButtons(system.indicatorTarget, system.buttons)
}

func BruteForceButtons(target []bool, buttons [][]int) int {
	numberOfCombinations := int(math.Pow(2, float64(len(buttons)))) + 1

	minFound := 0
	for n := range numberOfCombinations {
		buttonsToPress := GetButtonsForNumber(n, len(buttons))
		if minFound > 0 && minFound <= len(buttonsToPress) {
			continue
		}

		res := make([]bool, len(target))
		for _, buttonNr := range buttonsToPress {
			for _, i := range buttons[buttonNr] {
				res[i] = !res[i]
			}
		}

		if EqualBoolSlices(target, res) {
			minFound = len(buttonsToPress)
		}
	}

	return minFound
}

var lineToSystemRegex = regexp.MustCompile(`\[([\.#]+)\] \((.+)\) \{([\d,]+)`)

func LineToSystem(line string) *System {
	matches := lineToSystemRegex.FindStringSubmatch(line)

	return &System{
		indicatorTarget: StrToIndicatorTarget(matches[1]),
		buttons:         StrToButtons(matches[2]),
		joltage:         StrToInts(matches[3]),
	}
}

func StrToIndicatorTarget(str string) []bool {
	res := make([]bool, len(str))
	for i, char := range str {
		if char == '#' {
			res[i] = true
		}
	}
	return res
}

func StrToButtons(str string) [][]int {
	strParts := strings.Split(str, ") (")
	res := make([][]int, len(strParts))

	for i, s := range strParts {
		res[i] = StrToInts(s)
	}

	return res
}

func StrToInts(str string) []int {
	strParts := strings.Split(str, ",")
	res := make([]int, len(strParts))

	for i, s := range strParts {
		res[i], _ = strconv.Atoi(s)
	}

	return res
}

func GetButtonsForNumber(in int, size int) []int {
	res := []int{}

	for i := size; i > 0; i-- {
		if in%2 == 1 {
			res = append(res, i-1)
		}
		in = in / 2
	}

	return res
}

func EqualBoolSlices(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
