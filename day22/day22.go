package main

import (
	. "../util"
	"fmt"
	"strconv"
	"strings"
)

type DataType struct {
	player1 Deck
	player2 Deck
}

type Deck struct {
	elements []int
}

func parseData() DataType {
	data := FetchInputData(22)
	dataSplit := strings.Split(data, "\n\n")

	player1 := make([]int, 0)
	for _, v := range strings.Split(dataSplit[0], "\n")[1:] {
		vi, _ := strconv.Atoi(v)
		player1 = append(player1, vi)
	}

	player2 := make([]int, 0)
	for _, v := range strings.Split(dataSplit[1], "\n")[1:] {
		vi, _ := strconv.Atoi(v)
		player2 = append(player2, vi)
	}

	return DataType{Deck{player1}, Deck{player2}}
}

func (d *Deck) DrawACard() int {
	card := d.elements[0]
	d.elements = d.elements[1:]

	return card
}

func (d *Deck) PutOnBottom(card1 int, card2 int) {
	d.elements = append(d.elements, card1, card2)
}

func (d *Deck) Score() (result int) {
	for i := 0; i < len(d.elements); i++ {
		result += d.elements[(len(d.elements)-i-1)] * (i + 1)
	}

	return
}

func (d *Deck) String() string {
	var sb strings.Builder
	sb.Grow(len(d.elements) * 2)
	for _, v := range d.elements {
		_, _ = fmt.Fprintf(&sb, "%d", v)
	}
	return sb.String()
}

func (d *Deck) Copy(n int) Deck {
	newElements := make([]int, n)
	for i := 0; i < n; i++ {
		newElements[i] = d.elements[i]
	}

	return Deck{newElements}
}

func solvePart1(data DataType) (rc int) {
	for {
		if len(data.player1.elements) == 0 {
			return data.player2.Score()
		}

		if len(data.player2.elements) == 0 {
			return data.player1.Score()
		}

		card1 := data.player1.DrawACard()
		card2 := data.player2.DrawACard()

		if card1 > card2 {
			data.player1.PutOnBottom(card1, card2)
		} else {
			data.player2.PutOnBottom(card2, card1)
		}
	}
}

func playAGame(player1 *Deck, player2 *Deck) {
	cache := make(map[string]bool)
	for {
		if len(player1.elements) == 0 {
			return
		}

		if len(player2.elements) == 0 {
			return
		}

		hash := player1.String() + "|" + player2.String()
		if _, ok := cache[hash]; ok {
			player1.elements = make([]int, 1)
			player2.elements = make([]int, 0)
			return
		}
		cache[hash] = true

		card1 := player1.DrawACard()
		card2 := player2.DrawACard()

		if card1 <= len(player1.elements) && card2 <= len(player2.elements) {
			player1Copy := player1.Copy(card1)
			player2Copy := player2.Copy(card2)
			playAGame(&player1Copy, &player2Copy)

			if len(player1Copy.elements) > 0 {
				player1.PutOnBottom(card1, card2)
			} else {
				player2.PutOnBottom(card2, card1)
			}
		} else {
			if card1 > card2 {
				player1.PutOnBottom(card1, card2)
			} else {
				player2.PutOnBottom(card2, card1)
			}
		}
	}
}

func solvePart2(data DataType) (rc int) {
	playAGame(&data.player1, &data.player2)

	if len(data.player1.elements) == 0 {
		return data.player2.Score()
	} else {
		return data.player1.Score()
	}
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
