package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []int

func parseData() DataType {
	data := FetchInputData(1)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	for i, v1 := range data {
		for _, v2 := range data[i+1:] {
			if v1+v2 == 2020 {
				return v1 * v2
			}
		}
	}

	return
}

func solvePart2(data DataType) (rc int) {
	for i, v1 := range data {
		for j, v2 := range data[i+1:] {
			for _, v3 := range data[j+1:] {
				if v1+v2+v3 == 2020 {
					return v1 * v2 * v3
				}
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
