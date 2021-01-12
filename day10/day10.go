package main

import (
	. "../util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type DataType []int

func parseData() DataType {
	data := FetchInputData(10)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i], _ = strconv.Atoi(v)
	}

	sort.Ints(result)
	return result
}

func solvePart1(data DataType) (rc int) {
	differences := make(map[int]int)

	prev := 0
	for _, x := range data {
		differences[x-prev] += 1
		prev = x
	}

	return differences[1] * (differences[3] + 1)
}

func solvePart2(data DataType) (rc int) {
	data = append(data, data[len(data)-1]+3)

	lastUsed := map[int]int{0: 1}
	lastPrevUsed := lastUsed
	for i := 0; i < len(data)-1; i++ {
		lastPrevUsed = lastUsed
		lastUsed = make(map[int]int)
		for f, v := range lastPrevUsed {
			if data[i+1]-f <= 3 {
				lastUsed[f] += v
			}
			lastUsed[data[i]] += v
		}
	}

	for _, v := range lastUsed {
		return v
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
