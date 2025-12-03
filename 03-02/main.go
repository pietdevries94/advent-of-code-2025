package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(inputFile)

	total := 0

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		bank := inputScanner.Text()

		total += BatteryBankParser(bank)
	}

	// print the number of zeroes
	fmt.Println(total)
}

func BatteryBankParser(bank string) int {
	lastPickedPos := 0
	batteries := new(strings.Builder)

	num_of_batteries := 12

	for batNumber := range num_of_batteries {
		bankStart := lastPickedPos
		bankEnd := len(bank) - num_of_batteries + batNumber + 1

		bat := -1
		curPos := -1
		for pos, batteryRune := range bank[bankStart:bankEnd] {
			battery := int(batteryRune - '0')
			if battery > bat {
				bat = battery
				curPos = pos
			}
		}

		lastPickedPos += curPos + 1
		batteries.WriteString(strconv.Itoa(bat))
	}

	int, _ := strconv.Atoi(batteries.String())
	return int
}
