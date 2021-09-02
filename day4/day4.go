package main

import (
	. "../util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

type DataType struct {
	passports    []Passport
	requiredKeys []string
}

func parseData() DataType {
	data := FetchInputData(4)
	dataSplit := strings.Split(data, "\n\n")

	result := make([]Passport, len(dataSplit))
	for i, line := range dataSplit {
		tmp := make(Passport)
		for _, entry := range strings.Fields(line) {
			s := strings.Split(entry, ":")
			tmp[s[0]] = s[1]
		}
		result[i] = tmp
	}

	return DataType{result, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}}
}

func hasAllRequiredKeys(passport Passport, requiredKeys []string) bool {
	for _, key := range requiredKeys {
		if _, ok := passport[key]; !ok {
			return false
		}
	}

	return true
}

func isPassportEntryValid(key string, value string) bool {
	switch key {
	case "byr":
		v, _ := strconv.Atoi(value)
		return 1920 <= v && v <= 2002

	case "iyr":
		v, _ := strconv.Atoi(value)
		return 2010 <= v && v <= 2020

	case "eyr":
		v, _ := strconv.Atoi(value)
		return 2020 <= v && v <= 2030

	case "hgt":
		r := regexp.MustCompile(`^(\d+)(cm|in)$`)
		match := r.FindStringSubmatch(value)

		if match == nil {
			return false
		}

		v, _ := strconv.Atoi(match[1])
		if match[2] == "cm" {
			return 150 <= v && v <= 193
		} else {
			return 59 <= v && v <= 76
		}

	case "hcl":
		r := regexp.MustCompile(`^#(\d|[a-f]){6}$`)
		return r.MatchString(value)

	case "ecl":
		r := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
		return r.MatchString(value)

	case "pid":
		r := regexp.MustCompile(`^\d{9}$`)
		return r.MatchString(value)

	case "cid":
		return true
	}

	return false
}

func isPassportValid(passport Passport, requiredKeys []string) bool {
	if !hasAllRequiredKeys(passport, requiredKeys) {
		return false
	}

	for key, value := range passport {
		if !isPassportEntryValid(key, value) {
			return false
		}
	}

	return true
}

func solve(data DataType, validator func(Passport, []string) bool) (rc int) {
	for _, passport := range data.passports {
		if validator(passport, data.requiredKeys) {
			rc++
		}
	}

	return
}

func solvePart1(data DataType) (rc int) {
	return solve(data, hasAllRequiredKeys)
}

func solvePart2(data DataType) (rc int) {
	return solve(data, isPassportValid)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
