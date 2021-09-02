package main

import (
	. "../util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DataType []string

func parseData() DataType {
	data := FetchInputData(18)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i] = v
	}

	return result
}

func solveSimple(s string) int {
	sSplit := strings.Fields(s)
	result, _ := strconv.Atoi(sSplit[0])

	for i := 1; i < len(sSplit); i += 2 {
		operator := sSplit[i]
		nextNumber, _ := strconv.Atoi(sSplit[i + 1])
		switch operator {
		case "+":
			result += nextNumber
		case "*":
			result *= nextNumber
		}
	}

	return result
}

func solver(eq string, re *regexp.Regexp, eval func(string) int) int {
	for {
		if re.MatchString(eq) {
			eq = re.ReplaceAllStringFunc(eq, func (s string) string {
				return strconv.Itoa(eval(strings.Trim(s, "()")))
			})
		} else {
			return eval(eq)
		}
	}
}

func solvePart1(data DataType) (rc int) {
	r := regexp.MustCompile(`\([^()]+\)`)

	for _, eq := range data {
		rc += solver(eq, r, solveSimple)
	}

	return
}

func solvePart2(data DataType) (rc int) {
	r := regexp.MustCompile(`\([^()]+\)`)
	rPlus := regexp.MustCompile(`\d+ \+ \d+`)

	for _, eq := range data {
		rc += solver(eq, r, func(s string) int {
			return solver(s, rPlus, solveSimple)
		})
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
