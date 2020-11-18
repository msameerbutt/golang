package main

import (
	"fmt"
)

// Create a New type of 'deck'
// Which is a slice of strings
type deck [] string

// get new deck
func newDeck () deck {
	cards := deck {}
	cardSuits := [] string {"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := [] string {"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}
	return cards
}

// print deck of cards
func (d deck) print () {
	for i, card := range d {
		fmt.Println(i, "card:", card)
	}
}

// deal function
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}