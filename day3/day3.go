package main

import (
	. "../util"
	"fmt"
)

type DataType struct {
	grid Grid
}

func parseData() DataType {
	data := FetchInputData(3)
	grid := NewGridFromString(data, '#')
	return DataType{grid}
}

func findTrees(grid Grid, dx int, dy int) (rc int) {
	maxX, maxY := 0, 0
	for location := range grid {
		if maxX < location.X {
			maxX = location.X
		}

		if maxY < location.Y {
			maxY = location.Y
		}
	}

	x, y := 0, 0
	for y <= maxY {
		if grid.Contains(Location{X: x, Y: y}) {
			rc++
		}

		x = (x + dx) % (maxX + 1)
		y += dy
	}

	return
}

func solvePart1(data DataType) (rc int) {
	return findTrees(data.grid, 3, 1)
}

func solvePart2(data DataType) (rc int) {
	locations := []Location{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	rc = 1
	for _, location := range locations {
		rc *= findTrees(data.grid, location.X, location.Y)
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
