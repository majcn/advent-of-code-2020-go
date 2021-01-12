package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []string

func parseData() DataType {
	data := FetchInputData(18)
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i] = strings.ReplaceAll(v, " ", "")
	}

	return result
}

func findNumber(eq string, start int) (int, int) {
	i := start
	for i < len(eq) {
		x := eq[i]
		if '0' <= x && x <= '9' {
			i += 1
		} else {
			break
		}
	}

	number, _ := strconv.Atoi(eq[start:i])
	return number, i - start - 1
}

func solveSimplePart1(numbers []int, operators []byte) (result int) {
	result = numbers[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case '+':
			result = result + numbers[i+1]
			break
		case '*':
			result = result * numbers[i+1]
			break
		}
	}

	return
}

func solveSimplePart2(numbers []int, operators []byte) (result int) {
	i := 0
	for i < len(operators) {
		if operators[i] == '+' {
			left := numbers[i]
			right := numbers[i+1]

			operators = append(operators[:i], operators[i+1:]...)
			numbers = append(numbers[:i+1], numbers[i+2:]...)
			numbers[i] = left + right
		} else {
			i += 1
		}
	}

	return solveSimplePart1(numbers, operators)
}

func solver(eq string, start int, solveSimple func([]int, []byte) int) (int, int) {
	i := start

	operators := make([]byte, 0)
	numbers := make([]int, 0)
	for i < len(eq) {
		switch eq[i] {
		case '(':
			n, skipI := solver(eq, i+1, solveSimple)
			numbers = append(numbers, n)
			i += skipI
			break
		case ')':
			return solveSimple(numbers, operators), i - start + 1
		case '+':
			operators = append(operators, '+')
			break
		case '*':
			operators = append(operators, '*')
			break
		default:
			n, skipI := findNumber(eq, i)
			numbers = append(numbers, n)
			i += skipI
		}

		i += 1
	}

	return solveSimple(numbers, operators), 0
}

func solvePart1(data DataType) (rc int) {
	for _, eq := range data {
		el, _ := solver(eq, 0, solveSimplePart1)
		rc += el
	}

	return
}

func solvePart2(data DataType) (rc int) {
	for _, eq := range data {
		el, _ := solver(eq, 0, solveSimplePart2)
		rc += el
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
