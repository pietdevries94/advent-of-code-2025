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

	matrix := [][]int{}

	// When scanning a file, it does it line by line, making it easy to grabbing the instruction
	for inputScanner.Scan() {
		textLine := inputScanner.Text()
		line := []int{}

		for _, r := range textLine {
			if r == '@' {
				line = append(line, 1)
			} else {
				line = append(line, 0)
			}
		}
		matrix = append(matrix, line)
	}

	res := ParseMatrix(matrix, 0)
	fmt.Println(res)
}

func ParseMatrix(matrix [][]int, totalViableRolls int) int {
	minLineIndex := 0
	maxLineIndex := len(matrix) - 1
	minColIndex := 0
	maxColIndex := len(matrix[0]) - 1

	viableRolls := 0

	newMatrix := [][]int{}

	for lineIndex, line := range matrix {
		newLine := []int{}
		for colIndex, val := range line {
			if val == 0 {
				newLine = append(newLine, 0)
				continue
			}
			totalSurroundingRolls := 0

			// Get the rolls on this line
			if colIndex > minColIndex {
				totalSurroundingRolls += line[colIndex-1]
			}
			if colIndex < maxColIndex {
				totalSurroundingRolls += line[colIndex+1]
			}

			// Get the rolls on the previous line
			if lineIndex > minLineIndex {
				prevLine := matrix[lineIndex-1]
				if colIndex > minColIndex {
					totalSurroundingRolls += prevLine[colIndex-1]
				}
				totalSurroundingRolls += prevLine[colIndex]
				if colIndex < maxColIndex {
					totalSurroundingRolls += prevLine[colIndex+1]
				}
			}

			// Get the rolls on the previous line
			if lineIndex < maxLineIndex {
				nextLine := matrix[lineIndex+1]
				if colIndex > minColIndex {
					totalSurroundingRolls += nextLine[colIndex-1]
				}
				totalSurroundingRolls += nextLine[colIndex]
				if colIndex < maxColIndex {
					totalSurroundingRolls += nextLine[colIndex+1]
				}
			}

			if totalSurroundingRolls < 4 {
				viableRolls++
				newLine = append(newLine, 0)
			} else {
				newLine = append(newLine, 1)
			}
		}
		newMatrix = append(newMatrix, newLine)
	}

	totalViableRolls += viableRolls

	if viableRolls > 0 {
		return ParseMatrix(newMatrix, totalViableRolls)
	}

	return totalViableRolls
}
