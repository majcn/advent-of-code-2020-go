package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType struct {
	rules map[string][][]string
	codes []string
}

func parseData() DataType {
	data := FetchInputData(19)
	dataSplit := strings.Split(data, "\n\n")

	rules := make(map[string][][]string, len(dataSplit[0]))
	for _, v := range strings.Split(dataSplit[0], "\n") {
		ruleAsStrSplit := strings.Split(v, ": ")
		left := ruleAsStrSplit[0]

		right := ruleAsStrSplit[1]
		right = strings.ReplaceAll(right, "\"", "")
		rightAsSplit := strings.Split(right, " | ")

		rules[left] = make([][]string, len(rightAsSplit))
		for i, el := range strings.Split(right, " | ") {
			rules[left][i] = strings.Fields(el)
		}
	}

	codes := strings.Split(dataSplit[1], "\n")

	return DataType{rules, codes}
}

func checker(text string, key string, rules map[string][][]string) (result []string) {
	for _, rule := range rules[key] {
		partialResult := []string{""}

		if rule[0] == "a" || rule[0] == "b" {
			partialResultTmp := make([]string, 0, len(partialResult))
			for _, r := range partialResult {
				if strings.HasPrefix(text, r+rule[0]) {
					partialResultTmp = append(partialResultTmp, r+rule[0])
				}
			}
			partialResult = partialResultTmp
		} else {
			for _, rulePart := range rule {
				var partialResultTmp []string
				for _, r1 := range partialResult {
					cText := text[len(r1):]
					for _, r2 := range checker(cText, rulePart, rules) {
						if strings.HasPrefix(text, r1+r2) {
							partialResultTmp = append(partialResultTmp, r1+r2)
						}
					}
				}
				partialResult = partialResultTmp
			}
		}

		result = append(result, partialResult...)
	}

	return
}

func solve(data DataType) (rc int) {
	for _, c := range data.codes {
		for _, cc := range checker(c, "0", data.rules) {
			if c == cc {
				rc++
				break
			}
		}
	}

	return
}

func solvePart1(data DataType) (rc int) {
	return solve(data)
}

func solvePart2(data DataType) (rc int) {
	data.rules["8"] = [][]string{{"42"}, {"42", "8"}}
	data.rules["11"] = [][]string{{"42", "31"}, {"42", "11", "31"}}

	return solve(data)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
