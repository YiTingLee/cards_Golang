package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 2)

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}
