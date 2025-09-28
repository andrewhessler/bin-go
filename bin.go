package main

import (
	"fmt"
	"math/rand/v2"
)

var p = fmt.Println

var bNumbers = make([]int, 15)
var iNumbers = make([]int, 15)
var nNumbers = make([]int, 15)
var gNumbers = make([]int, 15)
var oNumbers = make([]int, 15)

func initializeBingoNumbers() {
	for i := 0; i < len(bNumbers); i++ {
		bNumbers[i] = i
		iNumbers[i] = i + 15
		nNumbers[i] = i + 30
		gNumbers[i] = i + 45
		oNumbers[i] = i + 60
	}
}

type bingoCard struct {
	values []bool
}

func newCard() *bingoCard {
	p := bingoCard{}
	p.values = make([]bool, 75)

	shuffleIntSlice(bNumbers)
	shuffleIntSlice(iNumbers)
	shuffleIntSlice(nNumbers)
	shuffleIntSlice(gNumbers)
	shuffleIntSlice(oNumbers)

	for i := 0; i < 5; i++ {
		p.values[bNumbers[i]] = true
		p.values[iNumbers[i]] = true
		if i != 4 {
			p.values[nNumbers[i]] = true
		}
		p.values[gNumbers[i]] = true
		p.values[oNumbers[i]] = true
	}

	return &p

}

func shuffleIntSlice(array []int) {
	for i := 0; i < len(array); i++ {
		randInt := rand.IntN(len(array))
		array[i], array[randInt] = array[randInt], array[i]
	}
}

func runSimulation(balls []int, card *bingoCard) []int {
	turnsTaken := make([]int, 75)
	runs := 50_000_000
	count := 0

	for i := range runs {
		shuffleIntSlice(balls)
		if i%10_000_000 == 0 {
			p(i)
		}
		count = 0
		for j := range len(balls) {
			if card.values[j] == true {
				count++
				if count >= 24 {
					turnsTaken[j]++
					break
				}
			}
		}
	}
	return turnsTaken
}

func main() {
	initializeBingoNumbers()
	card := newCard()
	balls := make([]int, 75)
	for i := range len(balls) {
		balls[i] = i
	}
	p(runSimulation(balls, card))
}
