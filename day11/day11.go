package main

import (
	. "../util"
	"fmt"
	"reflect"
	"strings"
)

var MAXX int
var MAXY int

const (
	StateEmptySeat    SeatState = 1 // iota
	StateOccupiedSeat SeatState = 2 // iota
)

type SeatState int

type DataType map[Location]SeatState

func parseData() DataType {
	data := FetchInputData(11)
	dataSplit := strings.Split(data, "\n")

	MAXY = len(dataSplit)
	MAXX = len(dataSplit[0])

	result := make(DataType)
	for y, dy := range dataSplit {
		for x, dx := range dy {
			if string(dx) == "L" {
				result[Location{X: x, Y: y}] = StateEmptySeat
			}
		}
	}

	return result
}

func findNeighboursOccupied(data DataType, location Location) (rc int) {
	for _, n := range GetNeighbours8() {
		if data[location.Add(n)] == StateOccupiedSeat {
			rc++
		}
	}

	return
}

func findNeighboursOccupiedInf(data DataType, location Location) (rc int) {
	for _, neighbour := range GetNeighbours8() {
		nl := location.Add(neighbour)
		for data[nl] != StateEmptySeat && 0 <= nl.X && nl.X <= MAXX && 0 <= nl.Y && nl.Y <= MAXY {
			if data[nl] == StateOccupiedSeat {
				rc++
				break
			}

			nl = nl.Add(neighbour)
		}
	}

	return
}

func applyRulesPart1(data DataType, location Location) SeatState {
	neighboursOccupied := findNeighboursOccupied(data, location)

	if data[location] == StateEmptySeat && neighboursOccupied == 0 {
		return StateOccupiedSeat
	}

	if data[location] == StateOccupiedSeat && neighboursOccupied >= 4 {
		return StateEmptySeat
	}

	return data[location]
}

func applyRulesPart2(data DataType, location Location) SeatState {
	neighboursOccupiedInf := findNeighboursOccupiedInf(data, location)

	if data[location] == StateEmptySeat && neighboursOccupiedInf == 0 {
		return StateOccupiedSeat
	}

	if data[location] == StateOccupiedSeat && neighboursOccupiedInf >= 5 {
		return StateEmptySeat
	}

	return data[location]
}

func transform(data DataType, applyRules func(DataType, Location) SeatState) DataType {
	newData := make(DataType, len(data))
	for location := range data {
		newData[location] = applyRules(data, location)
	}
	return newData
}

func countOccupied(data DataType) (rc int) {
	for _, v := range data {
		if v == StateOccupiedSeat {
			rc++
		}
	}

	return
}

func solve(data DataType, applyRules func(DataType, Location) SeatState) int {
	oldData := transform(data, applyRules)
	newData := transform(oldData, applyRules)

	for !reflect.DeepEqual(oldData, newData) {
		oldData = newData
		newData = transform(oldData, applyRules)
	}

	return countOccupied(newData)
}

func solvePart1(data DataType) (rc int) {
	return solve(data, applyRulesPart1)
}

func solvePart2(data DataType) (rc int) {
	return solve(data, applyRulesPart2)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
