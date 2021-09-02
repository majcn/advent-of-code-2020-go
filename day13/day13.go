package main

import (
	. "../util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type DataType struct {
	myTimestamp int
	buses       []Bus
}

type Bus struct {
	id              int
	departTimestamp int
}

func parseData() DataType {
	data := FetchInputData(13)
	dataSplit := strings.Split(data, "\n")

	myTimestamp, _ := strconv.Atoi(dataSplit[0])

	busesSplit := strings.Split(dataSplit[1], ",")
	buses := make([]Bus, 0, len(busesSplit))
	for i, v := range busesSplit {
		busId, err := strconv.Atoi(v)
		if err == nil {
			buses = append(buses, Bus{busId, i})
		}
	}

	return DataType{myTimestamp, buses}
}

func calculateWaitingTime(myTimestamp int, busId int) int {
	return busId - (myTimestamp % busId)
}

func solvePart1(data DataType) (rc int) {
	minBusId := data.buses[0].id
	minWaitingTime := math.MaxInt64

	for _, bus := range data.buses {
		waitingTime := calculateWaitingTime(data.myTimestamp, bus.id)
		if waitingTime < minWaitingTime {
			minWaitingTime = waitingTime
			minBusId = bus.id
		}
	}

	return minBusId * minWaitingTime
}

func solvePart2(data DataType) (rc int) {
	step := 1
	for _, bus := range data.buses {
		for (rc+bus.departTimestamp)%bus.id != 0 {
			rc += step
		}
		step *= bus.id
	}

	return
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
