package main

import (
	. "../util"
	"fmt"
	"regexp"
	"strings"
)

type DataType [][]string

func parseData() DataType {
	data := FetchInputData(24)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = regexp.MustCompile(`(nw|ne|e|se|sw|w)`).FindAllString(line, -1)
	}

	return result
}

var directionMap = map[string]Location{
	"nw": {-1, 1},
	"ne": {1, 1},
	"e":  {2, 0},
	"se": {1, -1},
	"sw": {-1, -1},
	"w":  {-2, 0},
}

func getLocation(directions []string) Location {
	result := Location{}
	for _, direction := range directions {
		result = result.Add(directionMap[direction])
	}

	return result
}

func generateGrid(data DataType) Grid {
	grid := make(Grid)
	for _, directions := range data {
		l := getLocation(directions)
		if grid.Contains(l) {
			grid.Remove(l)
		} else {
			grid.Add(l)
		}
	}

	return grid
}

func solvePart1(data DataType) (rc int) {
	return len(generateGrid(data))
}

func solvePart2(data DataType) (rc int) {
	blackGrid := generateGrid(data)

	for i := 0; i < 100; i++ {
		whiteGrid := make(Grid)
		for l := range blackGrid {
			for _, a := range directionMap {
				adjacentLocation := l.Add(a)
				if !blackGrid.Contains(adjacentLocation) {
					whiteGrid.Add(adjacentLocation)
				}
			}
		}

		newBlackGrid := make(Grid)

		for l := range blackGrid {
			numberOfAdjacentBlackTiles := 0
			for _, a := range directionMap {
				adjacentLocation := l.Add(a)
				if blackGrid.Contains(adjacentLocation) {
					numberOfAdjacentBlackTiles++
				}
			}

			if numberOfAdjacentBlackTiles == 1 || numberOfAdjacentBlackTiles == 2 {
				newBlackGrid.Add(l)
			}
		}

		for l := range whiteGrid {
			numberOfAdjacentBlackTiles := 0
			for _, a := range directionMap {
				adjacentLocation := l.Add(a)
				if blackGrid.Contains(adjacentLocation) {
					numberOfAdjacentBlackTiles++
				}
			}

			if numberOfAdjacentBlackTiles == 2 {
				newBlackGrid.Add(l)
			}
		}

		blackGrid = newBlackGrid
	}

	return len(blackGrid)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
