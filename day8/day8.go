package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType []Command

func parseData() DataType {
	data := FetchInputData(8)
	dataSplit := strings.Split(data, "\n")

	result := make([]Command, len(dataSplit))
	for i, v := range dataSplit {
		_, _ = fmt.Sscanf(v, "%s %d", &result[i].cmd, &result[i].value)
	}
	return result
}

func solvePart1(data DataType) (rc int) {
	p := NewProgram(data)
	rc, _ = p.run()

	return
}

func solvePart2(data DataType) (rc int) {
	replacer := strings.NewReplacer("jmp", "nop", "nop", "jmp")

	for i := 0; i < len(data); i++ {
		cdata := make(DataType, len(data))
		copy(cdata, data)
		cdata[i].cmd = replacer.Replace(cdata[i].cmd)

		p := NewProgram(cdata)
		if result, err := p.run(); err == nil {
			return result
		}
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
