package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []Instruction

type Instruction struct {
	action byte
	value  int
}

func parseData() DataType {
	data := FetchInputData(12)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		action := v[0]
		value, _ := strconv.Atoi(v[1:])

		result[i] = Instruction{action, value}
	}

	return result
}

func nextLocation(l Location, code byte, value int) Location {
	switch code {
	case 'N':
		return l.Add(Location{Y: value})
	case 'E':
		return l.Add(Location{X: value})
	case 'S':
		return l.Add(Location{Y: -value})
	case 'W':
		return l.Add(Location{X: -value})
	}

	return Location{}
}

func nextDirectionPart1(direction byte, code byte, value int) (n byte) {
	value = value % 360
	if code == 'L' {
		value = 360 - value
	}

	n = direction
	for value > 0 {
		switch n {
		case 'N':
			n = 'E'
			break
		case 'E':
			n = 'S'
			break
		case 'S':
			n = 'W'
			break
		case 'W':
			n = 'N'
			break
		}
		value -= 90
	}

	return
}

func nextWaypointRotatePart2(waypoint Location, code byte, value int) (n Location) {
	value = value % 360
	if code == 'L' {
		value = 360 - value
	}

	n = waypoint
	for value > 0 {
		n = Location{X: n.Y, Y: -n.X}
		value -= 90
	}

	return
}

func solvePart1(data DataType) (rc int) {
	direction := byte('E')
	position := Location{}

	for _, instruction := range data {
		if instruction.action == 'R' || instruction.action == 'L' {
			direction = nextDirectionPart1(direction, instruction.action, instruction.value)
		} else if instruction.action == 'F' {
			position = nextLocation(position, direction, instruction.value)
		} else {
			position = nextLocation(position, instruction.action, instruction.value)
		}
	}

	return Abs(position.X) + Abs(position.Y)
}

func solvePart2(data DataType) (rc int) {
	waypoint := Location{X: 10, Y: 1}
	position := Location{}

	for _, instruction := range data {
		if instruction.action == 'R' || instruction.action == 'L' {
			waypoint = nextWaypointRotatePart2(waypoint, instruction.action, instruction.value)
		} else if instruction.action == 'F' {
			position = position.Add(waypoint.Mul(instruction.value))
		} else {
			waypoint = nextLocation(waypoint, instruction.action, instruction.value)
		}
	}

	return Abs(position.X) + Abs(position.Y)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
