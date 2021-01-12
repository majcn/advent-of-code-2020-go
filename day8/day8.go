package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []Command

func parseData() DataType {
	data := FetchInputData(8)
	dataSplit := strings.Split(data, "\n")

	result := make([]Command, len(dataSplit))
	for i := 0; i < len(dataSplit); i++ {
		fields := strings.Fields(dataSplit[i])
		command := new(Command)
		command.cmd = fields[0]
		command.value, _ = strconv.Atoi(fields[1])
		result[i] = *command
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	p := Program{make(map[int]bool), data, 0, 0, 0}

	for {
		p.step()
		if p.isInLoop() {
			return p.accumulator
		}
	}
}

func solvePart2(data DataType) (rc int) {
	changes := [][]string{
		{"jmp", "nop"},
		{"nop", "jmp"},
	}

	for _, change := range changes {
		changeFrom, changeTo := change[0], change[1]
		for i := 0; i < len(data); i++ {
			if data[i].cmd != changeFrom {
				continue
			}

			cdata := make(DataType, len(data))
			copy(cdata, data)
			cdata[i].cmd = changeTo

			p := Program{make(map[int]bool), cdata, 0, 0, 0}
			for {
				p.step()
				if p.isFinished() {
					return p.accumulator
				}
				if p.isInLoop() {
					break
				}
			}
		}
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
