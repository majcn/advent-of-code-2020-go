package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []int

func parseData() DataType {
	data := FetchInputData(15)
	dataSplit := strings.Split(data, ",")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func solve(data DataType, nrIterations int) int {
	lastElement := data[len(data)-1]
	prevLastElement := data[len(data)-2]

	value := make(map[int]int, len(data)-1)
	for i := 0; i < len(data)-1; i++ {
		value[data[i]] = i
	}

	for i := len(data) - 1; i < nrIterations; i++ {
		prevLastElement = lastElement
		if _, ok := value[lastElement]; ok {
			lastElement = i - value[lastElement]
		} else {
			lastElement = 0
		}
		value[prevLastElement] = i
	}

	return prevLastElement
}

func solvePart1(data DataType) (rc int) {
	return solve(data, 2020)
}

func solvePart2(data DataType) (rc int) {
	return solve(data, 30000000)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
