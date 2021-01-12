package main

import (
	. "../util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DataType map[string]map[string]string

func parseData() DataType {
	data := FetchInputData(7)
	dataSplit := strings.Split(data, "\n")

	r1 := regexp.MustCompile("^(.*) bags contain (.*)$")
	r2 := regexp.MustCompile("^([0-9]+) (.*) (bag|bags)\\.?$")

	result := make(DataType)
	for _, line := range dataSplit {
		match1 := r1.FindStringSubmatch(line)
		key, valueAsStr := match1[1], match1[2]

		value := make(map[string]string)
		for _, xx := range strings.Split(valueAsStr, ", ") {
			match2 := r2.FindStringSubmatch(xx)
			if match2 != nil {
				value[match2[2]] = match2[1]
			}
		}
		result[key] = value
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	var f func(string) int
	f = func(key string) int {
		if key == "shiny gold" {
			return 1
		}

		for vk := range data[key] {
			if f(vk) == 1 {
				return 1
			}
		}

		return 0
	}

	for x := range data {
		rc += f(x)
	}
	return rc - 1
}

func solvePart2(data DataType) (rc int) {
	var f func(string) int
	f = func(key string) int {
		_, ok := data[key]
		if !ok {
			return 1
		}

		result := 0
		for vk, vv := range data[key] {
			vvi, _ := strconv.Atoi(vv)
			result += f(vk) * vvi
		}
		return result + 1
	}

	return f("shiny gold") - 1
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
