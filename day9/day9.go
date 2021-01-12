package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []int

func parseData() DataType {
	data := FetchInputData(9)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func isNumberValid(part []int, number int) bool {
	for _, v1 := range part {
		for _, v2 := range part {
			if v1+v2 == number {
				return true
			}
		}
	}

	return false
}

func solvePart1(data DataType) (rc int) {
	preambleLen := 25

	for i := 0; i < len(data)-preambleLen; i++ {
		part := data[i : i+preambleLen]
		nextNumber := data[i+preambleLen]

		if !isNumberValid(part, nextNumber) {
			return nextNumber
		}
	}

	return
}

func solvePart2(data DataType) (rc int) {
	invalidNumber := solvePart1(data)

	for n := 2; n < len(data); n++ {
		for i := 0; i < len(data)-n; i++ {
			part := data[i : i+n]
			if Sum(part) == invalidNumber {
				return Min(part) + Max(part)
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
