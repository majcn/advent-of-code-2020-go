package main

import (
	. "../util"
	"strings"
)

type Tile struct {
	id   int
	data Grid
	size int
}

func (tile *Tile) getPixel(x int, y int) byte {
	if tile.data.Contains(Location{X: x, Y: y}) {
		return '#'
	} else {
		return '.'
	}
}

func (tile *Tile) rotate() {
	newData := make(Grid)

	for xy := range tile.data {
		newData.Add(Location{X: tile.size - xy.Y - 1, Y: xy.X})
	}

	tile.data = newData
}

func (tile *Tile) flip() {
	newData := make(Grid)

	for xy := range tile.data {
		newData.Add(Location{X: tile.size - xy.X - 1, Y: xy.Y})
	}

	tile.data = newData
}

func (tile *Tile) orientate(f func(*Tile) bool) {
	for i := 0; i < 8; i++ {
		if !f(tile) {
			tile.rotate()
			if i == 3 {
				tile.flip()
			}
		}
	}
}

func (tile *Tile) getLeft() string {
	var sb strings.Builder
	for y := 0; y < tile.size; y++ {
		sb.WriteByte(tile.getPixel(0, y))
	}

	return sb.String()
}

func (tile *Tile) getRight() string {
	var sb strings.Builder
	for y := 0; y < tile.size; y++ {
		sb.WriteByte(tile.getPixel(tile.size-1, y))
	}

	return sb.String()
}

func (tile *Tile) getBottom() string {
	var sb strings.Builder
	for x := 0; x < tile.size; x++ {
		sb.WriteByte(tile.getPixel(x, tile.size-1))
	}

	return sb.String()
}

func (tile *Tile) getTop() string {
	var sb strings.Builder
	for x := 0; x < tile.size; x++ {
		sb.WriteByte(tile.getPixel(x, 0))
	}

	return sb.String()
}

func (tile *Tile) getAllPossibleEdges() []string {
	left := tile.getLeft()
	right := tile.getRight()
	top := tile.getTop()
	bottom := tile.getBottom()

	return []string{
		left, ReverseString(left),
		right, ReverseString(right),
		top, ReverseString(top),
		bottom, ReverseString(bottom),
	}
}
