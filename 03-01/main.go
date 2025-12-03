package main

import (
	"bufio"
	"fmt"
	"os"
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
	firstBat := -1
	firstPos := -1

	// because if the last character is the largest, we need to 1 to largest as the first
	for pos, batteryRune := range bank[:len(bank)-1] {
		battery := int(batteryRune - '0')
		if battery > firstBat {
			firstBat = battery
			firstPos = pos
		}
	}

	secondBat := -1
	for _, batteryRune := range bank[firstPos+1:] {
		battery := int(batteryRune - '0')
		if battery > secondBat {
			secondBat = battery
		}
	}

	return firstBat*10 + secondBat
}
