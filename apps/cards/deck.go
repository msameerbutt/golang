package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
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
// this is a receiver function, it would be called like deck.print()
func (d deck) print () {
	for i, card := range d {
		fmt.Println(i, "card:", card)
	}
}

// deal function
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// will convert deck slices from string into a string
// a receiver function
func (d deck) toString () string {
	return strings.Join([]string(d), ",")
}

// will save a deck into a file on the disk
// a receiver function
func (d deck) saveToFile (filename string) error {
		return ioutil.WriteFile(filename, []byte (d.toString()), 0666)
}

// This is a helper function that will read content of a file
func newDeckFromFile (filename string) error {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println ("Error:", err)
		os.Exit(1)
	}
}