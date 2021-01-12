package main

import (
	. "../util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DataType struct {
	types         []Type
	myTicket      []int
	nearbyTickets [][]int
}

type Type struct {
	name string
	rule Rule
}

type Rule struct {
	from1 int
	to1   int
	from2 int
	to2   int
}

func parseData() DataType {
	data := FetchInputData(16)
	dataSplit := strings.Split(data, "\n\n")

	typesLineRe := regexp.MustCompile("^(.*): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$")
	typeLines := strings.Split(dataSplit[0], "\n")
	types := make([]Type, len(typeLines))
	for i, typesLine := range typeLines {
		match := typesLineRe.FindStringSubmatch(typesLine)
		name := match[1]
		ruleFrom1, _ := strconv.Atoi(match[2])
		ruleTo1, _ := strconv.Atoi(match[3])
		ruleFrom2, _ := strconv.Atoi(match[4])
		ruleTo2, _ := strconv.Atoi(match[5])

		types[i] = Type{name, Rule{ruleFrom1, ruleTo1, ruleFrom2, ruleTo2}}
	}

	myTicketLine := strings.Split(dataSplit[1], "\n")[1]
	myTicket := make([]int, 0)
	for _, v := range strings.Split(myTicketLine, ",") {
		vi, _ := strconv.Atoi(v)
		myTicket = append(myTicket, vi)
	}

	nearbyTicketsLines := strings.Split(dataSplit[2], "\n")[1:]
	nearbyTickets := make([][]int, len(nearbyTicketsLines))
	for i, nearbyTicketsLine := range nearbyTicketsLines {
		tmp := make([]int, 0)
		for _, v := range strings.Split(nearbyTicketsLine, ",") {
			vi, _ := strconv.Atoi(v)
			tmp = append(tmp, vi)
		}
		nearbyTickets[i] = tmp
	}

	return DataType{types, myTicket, nearbyTickets}
}

func applyRule(x int, rule Rule) bool {
	return (rule.from1 <= x && x <= rule.to1) || (rule.from2 <= x && x <= rule.to2)
}

func getInvalidFields(ticket []int, types []Type) []int {
	result := make([]int, 0, len(ticket))
	for _, t := range ticket {
		isValidField := false
		for _, fieldType := range types {
			if applyRule(t, fieldType.rule) {
				isValidField = true
				break
			}
		}

		if !isValidField {
			result = append(result, t)
		}
	}
	return result
}

func solvePart1(data DataType) (rc int) {
	for _, ticket := range data.nearbyTickets {
		invalidFields := getInvalidFields(ticket, data.types)
		rc += Sum(invalidFields)
	}
	return
}

func solvePart2(data DataType) (rc int) {
	validTickets := make([][]int, 0, len(data.nearbyTickets))
	for _, ticket := range data.nearbyTickets {
		invalidFields := getInvalidFields(ticket, data.types)
		if len(invalidFields) == 0 {
			validTickets = append(validTickets, ticket)
		}
	}

	validTypesForColumn := make([]StringSet, 0)
	for i := 0; i < len(validTickets[0]); i++ {
		validTypeNames := make(StringSet)
		for _, fieldType := range data.types {
			isValidFieldType := true
			for _, ticket := range validTickets {
				if !applyRule(ticket[i], fieldType.rule) {
					isValidFieldType = false
					break
				}
			}

			if isValidFieldType {
				validTypeNames.Add(fieldType.name)
			}
		}

		validTypesForColumn = append(validTypesForColumn, validTypeNames)
	}

	allOptions := make(StringSet)
	for _, fieldType := range data.types {
		allOptions.Add(fieldType.name)
	}

	fields := make(map[int]string)
	for len(allOptions) > 0 {
		for column, validTypes := range validTypesForColumn {
			if len(validTypes) == 1 {
				v := validTypes.Pop()
				fields[column] = v
				allOptions.Remove(v)
			}
		}

		validTypesForColumnTmp := make([]StringSet, len(validTypesForColumn))
		for i, validTypes := range validTypesForColumn {
			validTypesForColumnTmp[i] = validTypes.Intersection(&allOptions)
		}
		validTypesForColumn = validTypesForColumnTmp
	}

	rc = 1
	for i, name := range fields {
		if strings.HasPrefix(name, "departure") {
			rc *= data.myTicket[i]
		}
	}
	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
