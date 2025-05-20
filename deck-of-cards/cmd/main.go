package main

import (
	"fmt"

	"github.com/FedericoBarberon/Go-Exercises/deck"
)

func main() {
	fmt.Println(deck.New(deck.Deck(3), deck.Filter(func(c deck.Card) bool { return c.Suit == deck.Spades })))
}
