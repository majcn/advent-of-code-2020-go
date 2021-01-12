package main

import (
	. "../util"
	"fmt"
	"regexp"
	"strconv"
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

	r := regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		match := r.FindStringSubmatch(v)

		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		letter := match[3][0]
		password := match[4]

		result[i] = Entry{min, max, letter, password}
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	for _, entry := range data {
		tmp := 0
		for i := 0; i < len(entry.Password); i++ {
			if entry.Password[i] == entry.Letter {
				tmp++
			}
		}

		if entry.Min >= tmp && tmp <= entry.Max {
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
