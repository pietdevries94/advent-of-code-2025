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

	position := 50
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

		if negative {
			position -= steps
		} else {
			position += steps
		}

		if (position % 100) == 0 {
			numberOfZeroes++
		}
	}

	// print the number of zeroes
	fmt.Println(numberOfZeroes)
}
