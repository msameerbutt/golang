package main
 
func main () {
	// Decare and initlized cards array of string
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
} 