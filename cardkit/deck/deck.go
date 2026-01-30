package deck

import (
	"math/rand"
	"time"

	"github.com/cshabsin/thegrid/cardkit/card"
)

type Deck []*card.Card

func NewStandard52() Deck {
	deck := make(Deck, 0, 52)
	for suit := card.Clubs; suit <= card.Spades; suit++ {
		for rank := card.Ace; rank <= card.King; rank++ {
			deck = append(deck, &card.Card{Suit: suit, Rank: rank})
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

func (d *Deck) Draw() *card.Card {
	if len(*d) == 0 {
		return nil
	}
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}
