package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []int

func parseData() DataType {
	data := FetchInputData(25)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func calculateNextVal(val int, subjectNumber int) int {
	return (val * subjectNumber) % 20201227
}

func findLoopSize(goal int) int {
	subjectNumber := 7

	val := 1
	loopSize := 1
	for val != goal {
		val = calculateNextVal(val, subjectNumber)
		loopSize++
	}

	return loopSize - 1
}

func solvePart1(data DataType) (rc int) {
	cardPublicKey, doorPublicKey := data[0], data[1]

	loopSize := findLoopSize(cardPublicKey)
	rc = 1
	for i := 0; i < loopSize; i++ {
		rc = calculateNextVal(rc, doorPublicKey)
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
}
