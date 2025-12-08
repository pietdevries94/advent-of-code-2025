package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coords = struct {
	x, y, z int
}

type Pair = struct {
	a, b string
	dist float64
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputScanner := bufio.NewScanner(inputFile)

	var allCoords []string

	for inputScanner.Scan() {
		line := inputScanner.Text()
		allCoords = append(allCoords, line)
	}

	pairs := CreatePairs(allCoords)
	slices.SortFunc(pairs, func(a, b *Pair) int {
		return int(a.dist - b.dist)
	})

	coordGroups := map[string]int{}
	groupSize := map[int]int{}
	nextGroup := 1

	var pair *Pair

	i := 0
	for {
		if len(groupSize) == 1 && len(coordGroups) == len(allCoords) {
			break
		}

		pair = pairs[i]
		i++

		if coordGroups[pair.a] > 0 && coordGroups[pair.b] > 0 {
			if coordGroups[pair.a] == coordGroups[pair.b] {
				continue
			}

			// merge the groups
			targetGroup := coordGroups[pair.a]
			sourceGroup := coordGroups[pair.b]
			for k, group := range coordGroups {
				if group == sourceGroup {
					coordGroups[k] = targetGroup
				}
			}

			groupSize[targetGroup] += groupSize[sourceGroup]
			delete(groupSize, sourceGroup)
			continue
		}

		if group, ok := coordGroups[pair.a]; ok {
			coordGroups[pair.b] = group
			groupSize[group]++
			continue
		}
		if group, ok := coordGroups[pair.b]; ok {
			coordGroups[pair.a] = group
			groupSize[group]++
			continue
		}
		coordGroups[pair.a] = nextGroup
		coordGroups[pair.b] = nextGroup
		groupSize[nextGroup] = 2
		nextGroup++
	}

	res := StringToCoords(pair.a).x * StringToCoords(pair.b).x

	fmt.Println(res)
}

func CreatePairs(coords []string) []*Pair {
	pairs := []*Pair{}
	seen := map[string]map[string]bool{}

	for _, c1 := range coords {
		seen[c1] = map[string]bool{}
		for _, c2 := range coords {
			if c1 == c2 {
				continue
			}
			if _, ok := seen[c1][c2]; ok {
				continue
			}
			if _, ok := seen[c2][c1]; ok {
				continue
			}

			dist := CalculateDistance(c1, c2)

			seen[c1][c2] = true
			pairs = append(pairs, &Pair{a: c1, b: c2, dist: dist})
		}
	}

	return pairs
}

func StringToCoords(str string) Coords {
	parts := strings.Split(str, ",")
	c := Coords{}
	c.x, _ = strconv.Atoi(parts[0])
	c.y, _ = strconv.Atoi(parts[1])
	c.z, _ = strconv.Atoi(parts[2])
	return c
}

func CalculateDistance(aStr, bStr string) float64 {
	a := StringToCoords(aStr)
	b := StringToCoords(bStr)
	x := GetDiff(a.x, b.x)
	y := GetDiff(a.y, b.y)
	z := GetDiff(a.z, b.z)

	dPow2 := math.Pow(float64(x), 2) + math.Pow(float64(y), 2) + math.Pow(float64(z), 2)
	return math.Sqrt(dPow2)
}

func GetDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
