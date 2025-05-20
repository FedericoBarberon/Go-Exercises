package deck

import "fmt"

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker //Special
)

var suits = [...]Suit{Spades, Diamonds, Clubs, Hearts}

type Card struct {
	Rank
	Suit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

func (c Rank) String() string {
	switch c {
	case Ace:
		return "Ace"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("%d", c)
	}
}

func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Hearts:
		return "Hearts"
	case Joker:
		return "Joker"
	default:
		return fmt.Sprintf("%d", s)
	}
}
