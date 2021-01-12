package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType []Tile

func parseData() DataType {
	data := FetchInputData(20)
	dataSplit := strings.Split(data, "\n\n")

	result := make(DataType, len(dataSplit))
	for i, tileAsStr := range dataSplit {
		tileAsStrSplit := strings.SplitN(tileAsStr, "\n", 2)
		tileId, _ := strconv.Atoi(tileAsStrSplit[0][5 : len(tileAsStrSplit[0])-1])
		grid := NewGridFromString(tileAsStrSplit[1], '#')
		result[i] = Tile{tileId, grid, strings.Index(tileAsStrSplit[1], "\n")}
	}

	return result
}

func findOneCorner(availableTiles map[int]*Tile) (*Tile, *Tile, *Tile) {
	tmpResult := [2]*Tile{nil, nil}

loop:
	for _, t1 := range availableTiles {
		nCounter := 0
		for _, t2 := range availableTiles {
			if t1 == t2 {
				continue
			}

			if !Disjoint(t1.getAllPossibleEdges(), t2.getAllPossibleEdges()) {
				if nCounter == 2 {
					continue loop
				}

				tmpResult[nCounter] = t2
				nCounter++
			}
		}

		if nCounter == 2 {
			return t1, tmpResult[0], tmpResult[1]
		}
	}

	return nil, nil, nil
}

func orientateTopLeftCorner(corner *Tile, n1 *Tile, n2 *Tile) {
	n1AllPossibleEdges := n1.getAllPossibleEdges()
	n2AllPossibleEdges := n2.getAllPossibleEdges()

	corner.orientate(func(item *Tile) bool {
		m1 := !Disjoint([]string{item.getRight()}, n1AllPossibleEdges)
		m2 := !Disjoint([]string{item.getBottom()}, n2AllPossibleEdges)
		return m1 && m2
	})
}

func findNextTileRight(current *Tile, availableTiles map[int]*Tile) *Tile {
	currentRight := current.getRight()
	for _, tile := range availableTiles {
		for _, edge := range tile.getAllPossibleEdges() {
			if currentRight == edge {
				return tile
			}
		}
	}

	return nil
}

func findNextTileBottom(current *Tile, availableTiles map[int]*Tile) *Tile {
	currentRight := current.getBottom()
	for _, tile := range availableTiles {
		for _, edge := range tile.getAllPossibleEdges() {
			if currentRight == edge {
				return tile
			}
		}
	}

	return nil
}

func orientateNextTileRight(current *Tile, nextTile *Tile) {
	currentRight := current.getRight()
	nextTile.orientate(func(item *Tile) bool {
		return currentRight == item.getLeft()
	})
}

func orientateNextTileBottom(current *Tile, nextTile *Tile) {
	currentBottom := current.getBottom()
	nextTile.orientate(func(item *Tile) bool {
		return currentBottom == item.getTop()
	})
}

func getImage(data []Tile) [][]Tile {
	availableTiles := make(map[int]*Tile, len(data))
	for i := range data {
		availableTiles[data[i].id] = &data[i]
	}

	corner, n1, n2 := findOneCorner(availableTiles)
	orientateTopLeftCorner(corner, n1, n2)
	delete(availableTiles, corner.id)

	result := make([][]Tile, 0)

	mostLeftItem := corner
	for mostLeftItem != nil {
		tmpResult := []Tile{*mostLeftItem}

		item := mostLeftItem
		for item != nil {
			nextTile := findNextTileRight(item, availableTiles)
			if nextTile != nil {
				orientateNextTileRight(item, nextTile)
				delete(availableTiles, nextTile.id)
				tmpResult = append(tmpResult, *nextTile)
			}
			item = nextTile
		}

		nextMostLeftItem := findNextTileBottom(mostLeftItem, availableTiles)
		if nextMostLeftItem != nil {
			orientateNextTileBottom(mostLeftItem, nextMostLeftItem)
			delete(availableTiles, nextMostLeftItem.id)
		}
		mostLeftItem = nextMostLeftItem

		result = append(result, tmpResult)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	image := getImage(data)
	imageSize := len(image)

	c1 := image[0][0]
	c2 := image[0][imageSize-1]
	c3 := image[imageSize-1][0]
	c4 := image[imageSize-1][imageSize-1]

	return c1.id * c2.id * c3.id * c4.id
}

func getNumberOfSeaMonsters(tile Tile, seaMonster []Location) int {
	numberOfSeaMonsters := 0

loop:
	for xy := range tile.data {
		for _, seaMonsterElement := range seaMonster {
			if !tile.data.Contains(xy.Add(seaMonsterElement)) {
				continue loop
			}
		}

		numberOfSeaMonsters++
	}

	return numberOfSeaMonsters
}

func getSeaMonsters(bigTile Tile) (int, int) {
	seaMonster := []Location{
		{18, -1},
		{0, 0}, {5, 0}, {6, 0}, {11, 0}, {12, 0}, {17, 0}, {18, 0}, {19, 0},
		{1, 1}, {4, 1}, {7, 1}, {10, 1}, {13, 1}, {16, 1},
	}

	bigTile.orientate(func(tile *Tile) bool {
		return getNumberOfSeaMonsters(*tile, seaMonster) > 0
	})

	return getNumberOfSeaMonsters(bigTile, seaMonster), len(seaMonster)
}

func solvePart2(data DataType) (rc int) {
	image := getImage(data)
	imageSize := len(image)

	bigImageGrid := make(Grid)
	for imageY := 0; imageY < imageSize; imageY++ {
		for imageX := 0; imageX < imageSize; imageX++ {
			tile := image[imageY][imageX]
			for tileXY := range tile.data {
				if 0 < tileXY.X && tileXY.X < tile.size-1 && 0 < tileXY.Y && tileXY.Y < tile.size-1 {
					bigImageGrid.Add(Location{X: imageX, Y: imageY}.Mul(tile.size - 2).Add(tileXY))
				}
			}
		}
	}

	bigTile := Tile{-1, bigImageGrid, imageSize * (image[0][0].size - 2)}
	numberOfSeaMonsters, seaMonsterSize := getSeaMonsters(bigTile)

	return len(bigTile.data) - numberOfSeaMonsters*seaMonsterSize
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
