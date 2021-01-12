package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType [][]string

func parseData() DataType {
	data := FetchInputData(14)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i] = strings.Split(v, " = ")
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	mask := make(map[int]int)
	mem := make(map[int]int)

	for _, line := range data {
		code, value := line[0], line[1]
		if code == "mask" {
			mask = make(map[int]int)
			for i := 0; i < len(value); i++ {
				if value[i] == '0' {
					mask[len(value)-i-1] = 0
				} else if value[i] == '1' {
					mask[len(value)-i-1] = 1
				}
			}
		} else {
			memLoc, _ := strconv.Atoi(code[4 : len(code)-1])
			v, _ := strconv.Atoi(value)
			for mk, mv := range mask {
				if mv == 0 {
					v = v &^ (1 << mk)
				} else {
					v = v | (1 << mk)
				}
			}
			mem[memLoc] = v
		}
	}

	for _, v := range mem {
		rc += v
	}

	return
}

func getAllMemoryLocations(prev string, now string) []string {
	result := make([]string, 0)
	foundX := false

	for i := 0; i < len(now); i++ {
		if now[i] == 'X' {
			result = append(result, getAllMemoryLocations(prev+now[:i]+"0", now[(i+1):])...)
			result = append(result, getAllMemoryLocations(prev+now[:i]+"1", now[(i+1):])...)
			foundX = true
			break
		}
	}

	if !foundX {
		result = append(result, prev+now)
	}

	return result
}

func solvePart2(data DataType) (rc int) {
	masksStr := ""
	mem := make(map[int]int)

	for _, line := range data {
		code := line[0]
		value := line[1]

		if code == "mask" {
			masksStr = value
		} else {
			memLoc, _ := strconv.Atoi(code[4 : len(code)-1])
			v, _ := strconv.Atoi(value)

			binMemLocWithoutZeros := strconv.FormatInt(int64(memLoc), 2)
			binMemLoc := strings.Repeat("0", len(masksStr)-len(binMemLocWithoutZeros)) + binMemLocWithoutZeros
			mask := ""
			for i := 0; i < len(masksStr); i++ {
				if masksStr[i] == '0' {
					mask += string(binMemLoc[i])
				} else {
					mask += string(masksStr[i])
				}
			}

			for _, memoryLocation := range getAllMemoryLocations("", mask) {
				memoryLocationAsInt, _ := strconv.ParseInt(memoryLocation, 2, 64)
				mem[int(memoryLocationAsInt)] = v
			}
		}
	}

	for _, v := range mem {
		rc += v
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
