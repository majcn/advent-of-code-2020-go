package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType []Entry

type Entry struct {
	Min      int
	Max      int
	Letter   byte
	Password string
}

func parseData() DataType {
	data := FetchInputData(2)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		_, _ = fmt.Sscanf(v, "%d-%d %c: %s", &result[i].Min, &result[i].Max, &result[i].Letter, &result[i].Password)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	for _, entry := range data {
		tmp := strings.Count(entry.Password, string(entry.Letter))
		if entry.Min <= tmp && tmp <= entry.Max {
			rc++
		}
	}

	return
}

func solvePart2(data DataType) (rc int) {
	for _, entry := range data {
		firstChar := entry.Password[entry.Min-1]
		secondChar := entry.Password[entry.Max-1]

		if (firstChar == entry.Letter && secondChar != entry.Letter) || (firstChar != entry.Letter && secondChar == entry.Letter) {
			rc++
		}
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
