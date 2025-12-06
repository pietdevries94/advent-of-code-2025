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

	columnRunes := [][]rune{}
	columnOperators := []string{}

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		line := inputScanner.Text()

		firstChar := string(strings.TrimSpace(line)[0])
		if firstChar == "+" || firstChar == "*" {
			columnOperators = splitOnAnyNumberOfSpaces(line)
			break
		}

		for i, r := range line {
			if i > len(columnRunes)-1 {
				columnRunes = append(columnRunes, []rune{r})
			} else {
				columnRunes[i] = append(columnRunes[i], r)
			}
		}
	}

	columnNumbers := [][]int{[]int{}}
	columnIndex := 0
	for _, runes := range columnRunes {
		str := strings.TrimSpace(string(runes))
		if str == "" {
			columnIndex++
			columnNumbers = append(columnNumbers, []int{})
			continue
		}
		num, _ := strconv.Atoi(str)
		columnNumbers[columnIndex] = append(columnNumbers[columnIndex], num)
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
