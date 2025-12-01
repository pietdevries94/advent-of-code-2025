package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(inputFile)

	// We put it on a crazy high number so we don't have to deal with negative numbers
	position := 10000000050
	numberOfZeroes := 0

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		instruction := inputScanner.Text()

		// We get the first character and determine if this change is positive or negative
		negative := string(instruction[0]) == "L"

		steps, err := strconv.Atoi(string(instruction[1:]))
		if err != nil {
			panic(err)
		}

		newPosition := position
		if negative {
			newPosition -= steps
		} else {
			newPosition += steps
		}

		differenceInCycles := int(float64(position)/100) - int(float64(newPosition)/100)
		if differenceInCycles < 0 {
			differenceInCycles = -differenceInCycles
		}

		// if the starting position is 0 and the instruction is going left, it will see a difference, but we already counted that 0, so we remove it
		if negative && (position%100) == 0 {
			differenceInCycles--
		}

		// if the ending position is 0 and the instruction is going left, it won't see the difference, so we add it
		if negative && (newPosition%100) == 0 {
			differenceInCycles++
		}

		fmt.Printf("Diff: %v - Start: %v - %v - End: %v\n", differenceInCycles, position, instruction, newPosition)

		numberOfZeroes += differenceInCycles
		position = newPosition
	}

	// print the number of zeroes
	fmt.Println(numberOfZeroes)
}
