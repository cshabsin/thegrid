package card

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	}
	return ""
}

func (s Suit) Color() string {
	switch s {
	case Clubs, Spades:
		return "black"
	case Diamonds, Hearts:
		return "red"
	}
	return ""
}

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

func (r Rank) String() string {
	switch r {
	case Ace:
		return "A"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return fmt.Sprintf("%d", r)
	}
}

type Card struct {
	Suit   Suit
	Rank   Rank
	FaceUp bool
}

type Deck []*Card

func NewStandard52() Deck {
	deck := make(Deck, 0, 52)
	for suit := Clubs; suit <= Spades; suit++ {
		for rank := Ace; rank <= King; rank++ {
			deck = append(deck, &Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d *Deck) Draw() *Card {
	if len(*d) == 0 {
		return nil
	}
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}
