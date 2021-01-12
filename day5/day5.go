package main

import (
	. "../util"
	"fmt"
	"math"
	"strings"
)

type DataType []string

func parseData() DataType {
	data := FetchInputData(5)
	dataSplit := strings.Split(data, "\n")

	return dataSplit
}

func getPosition(code string) Location {
	rl, rh := 0, 127

	for _, x := range code[:7] {
		fd := float64(rh-rl) / 2
		d := int(math.Round(fd))

		if x == 'F' {
			rh -= d
		} else {
			rl += d
		}
	}

	cl, ch := 0, 7
	for _, x := range code[7:] {
		fd := float64(ch-cl) / 2
		d := int(math.Round(fd))

		if x == 'L' {
			ch -= d
		} else {
			cl += d
		}
	}

	return Location{X: rh, Y: ch}
}

func getId(location Location) int {
	return location.X*8 + location.Y
}

func solvePart1(data DataType) (rc int) {
	for _, code := range data {
		id := getId(getPosition(code))
		if rc < id {
			rc = id
		}
	}

	return
}

func solvePart2(data DataType) (rc int) {
	positions := make(Grid)
	for _, code := range data {
		positions.Add(getPosition(code))
	}

	for row := 0; row < 128; row++ {
		loop:
		for column := 0; column < 8; column++ {
			location := Location{X: row, Y: column}

			if positions.Contains(location) {
				continue
			}

			for _, neighbour := range GetNeighbours4() {
				if !positions.Contains(location.Add(neighbour)) {
					continue loop
				}
			}

			return getId(location)
		}
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
