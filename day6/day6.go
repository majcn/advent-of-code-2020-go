package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType [][]string

func parseData() DataType {
	data := FetchInputData(6)
	dataSplit := strings.Split(data, "\n\n")

	result := make(DataType, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = strings.Fields(line)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	for _, group := range data {
		result := make(map[rune]bool)
		for _, person := range group {
			for _, question := range person {
				result[question] = true
			}
		}
		rc += len(result)
	}

	return
}

func solvePart2(data DataType) (rc int) {
	for _, group := range data {
		result := make(map[rune]int)
		for _, person := range group {
			for _, question := range person {
				result[question]++
			}
		}

		for _, v := range result {
			if len(group) == v {
				rc += 1
			}
		}
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
