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

	columnNumbers := [][]int{}
	columnOperators := []string{}

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		line := inputScanner.Text()

		firstChar := string(strings.TrimSpace(line)[0])
		if firstChar == "+" || firstChar == "*" {
			columnOperators = splitOnAnyNumberOfSpaces(line)
			break
		}

		for i, str := range splitOnAnyNumberOfSpaces(line) {
			num, _ := strconv.Atoi(str)
			if i > len(columnNumbers)-1 {
				columnNumbers = append(columnNumbers, []int{num})
			} else {
				columnNumbers[i] = append(columnNumbers[i], num)
			}
		}
	}

	grandTotal := 0
	for i, list := range columnNumbers {
		operator := columnOperators[i]
		grandTotal += calculate(list, operator)
	}

	fmt.Println(grandTotal)
}

func splitOnAnyNumberOfSpaces(s string) []string {
	split := strings.Split(s, " ")
	res := []string{}
	for _, part := range split {
		if part != "" {
			res = append(res, part)
		}
	}
	return res
}

func calculate(list []int, operator string) int {
	total := 0
	for i, num := range list {
		if i == 0 {
			total = num
			continue
		}
		if operator == "+" {
			total += num
		} else {
			total *= num
		}
	}
	return total
}
