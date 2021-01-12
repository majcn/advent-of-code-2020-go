package main

import (
	. "../util"
	"fmt"
	"math/big"
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

	result := DataType{
		myTimestamp: myTimestamp,
		buses:       buses,
	}

	return result
}

func ChineseRemainderTheorem(a []*big.Int, n []*big.Int) *big.Int {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p)
}

func calculateWaitingTime(myTimestamp int, busId int) int {
	return busId - (myTimestamp % busId)
}

func solvePart1(data DataType) (rc int) {
	minBusId := data.buses[0].id
	minWaitingTime := calculateWaitingTime(data.myTimestamp, minBusId)

	for _, bus := range data.buses {
		waitingTime := calculateWaitingTime(data.myTimestamp, bus.id)
		if waitingTime < minWaitingTime {
			minWaitingTime = waitingTime
			minBusId = bus.id
		}
	}

	return minBusId * minWaitingTime
}

func solvePart2(data DataType) (rc *big.Int) {
	a := make([]*big.Int, 0, len(data.buses))
	n := make([]*big.Int, 0, len(data.buses))
	for _, bus := range data.buses {
		a = append(a, big.NewInt(int64(-bus.departTimestamp)))
		n = append(n, big.NewInt(int64(bus.id)))
	}

	x := ChineseRemainderTheorem(a, n)
	return x
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
