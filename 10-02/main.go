package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
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

	i := 1
	startTime := time.Now()

	var totalSync sync.Map
	var wg sync.WaitGroup

	for inputScanner.Scan() {
		line := inputScanner.Text()

		key := i
		i++
		wg.Go(func() {
			result := ParseLine(line)
			totalSync.Store(key, result)
			fmt.Printf("line %v - result %v\n", key, result)
		})
	}

	wg.Wait()

	total := 0
	totalSync.Range(func(key, value any) bool {
		total += value.(int)
		return true
	})

	fmt.Printf("\nTotal runtime:%v\nanswer: %v\n", time.Since(startTime), total)
}

func ParseLine(line string) int {
	system := LineToSystem(line)
	count, found := BruteForceButtons(system.joltage, system.buttons, 0, 0, &map[string]bool{})
	if !found {
		panic(line)
	}
	return count
}

func BruteForceButtons(target []int, buttons [][]int, count, alreadyFound int, failedStates *map[string]bool) (int, bool) {
	// We don't care about this path because we found a better one
	if alreadyFound > 0 && alreadyFound < count {
		return -1, false
	}

	// And if it's in the failedStates we know we don't have to follow this path
	if (*failedStates)[fmt.Sprint(target)] {
		return -1, false
	}

	buttons = RemoveUnpressableButtons(target, buttons)
	if len(buttons) == 0 {
		return -1, false
	}

	// First we look for a button that is the only button to lower a certain target
	if button, amount := FindSafeButtonToPress(target, buttons); button != nil {
		for _, i := range button {
			target[i] -= amount
		}
		count += amount
		if TargetEmpty(target) {
			return count, true
		}
		return BruteForceButtons(target, buttons, count, alreadyFound, failedStates)
	}

	// Sadly we don't know a logical candidate, so we randomly try one
	for _, button := range buttons {
		newTarget := slices.Clone(target)
		for _, i := range button {
			newTarget[i]--
		}
		if TargetEmpty(newTarget) {
			return count + 1, true
		}
		res, ok := BruteForceButtons(newTarget, buttons, count+1, alreadyFound, failedStates)
		if ok && (alreadyFound == 0 || alreadyFound > res) {
			alreadyFound = res
		} else {
			(*failedStates)[fmt.Sprint(newTarget)] = true
		}
	}

	if alreadyFound > 0 {
		return alreadyFound, true
	}

	// Sadly this path doesn't work
	return -1, false
}

func FindSafeButtonToPress(target []int, buttons [][]int) ([]int, int) {
	buttonsWithNumber := map[int][]int{}
	for buttonNr, button := range buttons {
		for _, num := range button {
			buttonsWithNumber[num] = append(buttonsWithNumber[num], buttonNr)
		}
	}

	for _, buttonnrs := range buttonsWithNumber {
		if len(buttonnrs) != 1 {
			continue
		}
		// we found our button!
		button := buttons[buttonnrs[0]]
		minAmount := target[button[0]]
		for _, i := range button {
			c := target[i]
			if c < minAmount {
				minAmount = c
			}
		}
		return button, minAmount
	}

	return nil, 0
}

func RemoveUnpressableButtons(target []int, buttons [][]int) [][]int {
	cantPressIfIndex := map[int]bool{}
	for i, v := range target {
		if v == 0 {
			cantPressIfIndex[i] = true
		}
	}
	if len(cantPressIfIndex) == 0 {
		return buttons
	}
	newButtons := [][]int{}
buttonLoop:
	for _, button := range buttons {
		for _, num := range button {
			if cantPressIfIndex[num] {
				continue buttonLoop
			}
		}
		newButtons = append(newButtons, button)
	}
	return newButtons
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

func TargetEmpty(target []int) bool {
	for _, i := range target {
		if i > 0 {
			return false
		}
	}
	return true
}
