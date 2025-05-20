package deck

import (
	"math/rand"
	"sort"
)

type option func([]Card) []Card

func New(options ...option) []Card {
	deck := make([]Card, 0, 52)
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}

	for _, opt := range options {
		deck = opt(deck)
	}

	return deck
}

func CustomSort(fn func(c1, c2 Card) bool) option {
	return func(deck []Card) []Card {
		sort.Slice(deck, func(i, j int) bool {
			return fn(deck[i], deck[j])
		})

		return deck
	}
}

func DefaultSort() option {
	return CustomSort(func(c1, c2 Card) bool {
		return (c1.Suit < c2.Suit) || (c1.Suit == c2.Suit && c1.Rank < c2.Rank)
	})
}

func Shuffle() option {
	return func(deck []Card) []Card {
		rand.Shuffle(len(deck), func(i, j int) {
			deck[i], deck[j] = deck[j], deck[i]
		})

		return deck
	}
}

func Jokers(n int) option {
	return func(deck []Card) []Card {
		for i := 0; i < n; i++ {
			deck = append(deck, Card{Rank: Rank(i), Suit: Joker})
		}

		return deck
	}
}

func Filter(fn func(Card) bool) option {
	return func(deck []Card) []Card {
		var filteredDeck []Card
		for _, card := range deck {
			if fn(card) {
				filteredDeck = append(filteredDeck, card)
			}
		}

		return filteredDeck
	}
}

func Deck(n int) option {
	return func(deck []Card) []Card {
		var newDeck []Card
		for i := 0; i < n; i++ {
			newDeck = append(newDeck, deck...)
		}

		return newDeck
	}

}
