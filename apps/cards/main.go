package main


func main () {
	// Decare and initlized cards array of string
	cards1 := newDeck()
	cards1.saveToFile("my_cards")

	cards2 := newDeckFromFile("hello")

} 