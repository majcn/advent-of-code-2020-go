package main

import (
	. "../util"
	"fmt"
	"math"
	"strings"
)

type DataType map[LocationXYXW]bool

type LocationXYXW struct {
	x int
	y int
	z int
	w int
}

func parseData() DataType {
	data := FetchInputData(17)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType)
	for y, line := range dataSplit {
		for x, v := range line {
			if v == '#' {
				result[LocationXYXW{x, y, 0, 0}] = true
			}
		}
	}

	return result
}

func findInactiveNeighbors(data map[LocationXYXW]bool, el LocationXYXW, dimensions int) map[LocationXYXW]bool {
	result := make(map[LocationXYXW]bool)

	minDw, maxDw := 0, 0
	if dimensions == 4 {
		minDw = -1
		maxDw = 1
	}

	for dw := minDw; dw <= maxDw; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					newEl := LocationXYXW{el.x + dx, el.y + dy, el.z + dz, el.w + dw}
					if el != newEl {
						if _, ok := data[newEl]; !ok {
							result[newEl] = true
						}
					}
				}
			}
		}
	}

	return result
}

func solve(data DataType, dimensions int) int {
	maxNeighbors := int(math.Pow(3.0, float64(dimensions)) - 1)

	activeQueue := data
	for i := 0; i < 6; i++ {
		inactiveQueue := make(map[LocationXYXW]bool)
		newActiveQueue := make(map[LocationXYXW]bool)

		for el := range activeQueue {
			inactiveNeighbors := findInactiveNeighbors(activeQueue, el, dimensions)
			activeNeighborsLen := maxNeighbors - len(inactiveNeighbors)
			if activeNeighborsLen == 2 || activeNeighborsLen == 3 {
				newActiveQueue[el] = true
			}

			for inactiveNeighbor := range inactiveNeighbors {
				inactiveQueue[inactiveNeighbor] = true
			}
		}

		for el := range inactiveQueue {
			inactiveNeighbors := findInactiveNeighbors(activeQueue, el, dimensions)
			activeNeighborsLen := maxNeighbors - len(inactiveNeighbors)
			if activeNeighborsLen == 3 {
				newActiveQueue[el] = true
			}
		}

		activeQueue = newActiveQueue
	}

	return len(activeQueue)
}

func solvePart1(data DataType) (rc int) {
	return solve(data, 3)
}

func solvePart2(data DataType) (rc int) {
	return solve(data, 4)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
