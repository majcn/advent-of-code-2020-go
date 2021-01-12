package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type Cup struct {
	label int
	next  *Cup
}

type DataType struct {
	cups  map[int]*Cup
	first *Cup
	last  *Cup
}

func (data *DataType) Copy() *DataType {
	cups := make(map[int]*Cup, len(data.cups))
	originalCup := data.first

	first := &Cup{label: originalCup.label}
	cups[first.label] = first
	originalCup = originalCup.next

	cup := first
	for i := 1; i < len(data.cups); i++ {
		tmp := &Cup{label: originalCup.label}
		cups[tmp.label] = tmp
		originalCup = originalCup.next

		cup.next = tmp
		cup = tmp
	}
	cup.next = first

	return &DataType{cups, first, cup}
}

func parseData() DataType {
	data := FetchInputData(23)

	cups := make(map[int]*Cup, len(data))

	firstLabel, _ := strconv.Atoi(string(data[0]))
	firstCup := &Cup{label: firstLabel}
	cups[firstCup.label] = firstCup

	dataSplitAsInt := make([]int, len(data)-1)
	for i, v := range data[1:] {
		dataSplitAsInt[i], _ = strconv.Atoi(string(v))
	}
	lastCup := generateNodes(cups, firstCup, dataSplitAsInt)

	return DataType{cups, firstCup, lastCup}
}

func generateNodes(cups map[int]*Cup, firstCup *Cup, labels []int) *Cup {
	cup := firstCup
	for _, l := range labels {
		tmp := &Cup{label: l}
		cups[tmp.label] = tmp

		cup.next = tmp
		cup = tmp
	}
	cup.next = firstCup

	return cup
}

func solve(cups map[int]*Cup, firstCup *Cup, minLabel int, maxLabel int, N int) *Cup {
	currentCup := firstCup
	for i := 0; i < N; i++ {
		pickedCup1 := currentCup.next
		pickedCup2 := pickedCup1.next
		pickedCup3 := pickedCup2.next

		currentCup.next = pickedCup3.next
		destinationCupLabel := currentCup.label
		for {
			destinationCupLabel--
			if destinationCupLabel < minLabel {
				destinationCupLabel = maxLabel
			}

			if destinationCupLabel != pickedCup1.label && destinationCupLabel != pickedCup2.label && destinationCupLabel != pickedCup3.label {
				break
			}
		}

		destinationCup := cups[destinationCupLabel]
		pickedCup3.next = destinationCup.next
		destinationCup.next = pickedCup1

		currentCup = currentCup.next
	}

	return cups[1]
}

func solvePart1(data DataType) (rc string) {
	minLabel, maxLabel, N := 1, 9, 100
	data = *data.Copy()

	cup1 := solve(data.cups, data.first, minLabel, maxLabel, N)

	var sb strings.Builder
	cup := cup1.next
	for cup != cup1 {
		_, _ = fmt.Fprintf(&sb, "%d", cup.label)
		cup = cup.next
	}
	return sb.String()
}

func solvePart2(data DataType) (rc int) {
	minLabel, maxLabelFromData, maxLabel, N := 1, 9, 1000000, 10000000
	data = *data.Copy()

	additionalCups := make([]int, 0, maxLabel)
	for i := maxLabelFromData + 1; i <= maxLabel; i++ {
		additionalCups = append(additionalCups, i)
	}

	generateNodes(data.cups, data.last, additionalCups)
	data.cups[maxLabel].next = data.first

	cup1 := solve(data.cups, data.first, minLabel, maxLabel, N)
	return cup1.next.label * cup1.next.next.label
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
