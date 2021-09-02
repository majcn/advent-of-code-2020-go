package main

import (
	. "../util"
	"container/ring"
	"fmt"
)

type DataType []int

func parseData() DataType {
	data := FetchInputData(23)

	result := make([]int, len(data))
	for i, v := range data {
		result[i] = int(v - '0')
	}

	return result
}

func solve(data []int, minLabel int, maxLabel int, N int) *ring.Ring {
	cupsMap := make(map[int]*ring.Ring, maxLabel)
	cups := ring.New(maxLabel)

	for _, v := range data {
		cupsMap[v] = cups
		cups.Value = v
		cups = cups.Next()
	}

	for i := len(data) + 1; i <= maxLabel; i++ {
		cupsMap[i] = cups
		cups.Value = i
		cups = cups.Next()
	}

	currentCup := cups
	for i := 0; i < N; i++ {
		pick := currentCup.Unlink(3)
		destinationCupLabel := currentCup.Value.(int)

		for {
			destinationCupLabel = destinationCupLabel - 1
			if destinationCupLabel < minLabel {
				destinationCupLabel = maxLabel
			}

			shouldBreak := true
			pick.Do(func(i interface{}) {
				if destinationCupLabel == i.(int) {
					shouldBreak = false
				}
			})
			if shouldBreak {
				break
			}
		}

		cupsMap[destinationCupLabel].Link(pick)
		currentCup = currentCup.Next()
	}

	return cupsMap[1]
}

func solvePart1(data DataType) (rc int) {
	minLabel, maxLabel, N := 1, 9, 100

	r := solve(data, minLabel, maxLabel, N)
	r.Unlink(r.Len() - 1).Do(func(i interface{}) {
		rc = rc * 10 + i.(int)
	})
	return rc
}

func solvePart2(data DataType) (rc int) {
	minLabel, maxLabel, N := 1, 1000000, 10000000

	r := solve(data, minLabel, maxLabel, N)
	return r.Next().Value.(int) * r.Move(2).Value.(int)
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
