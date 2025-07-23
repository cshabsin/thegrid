package pile

import (
	"math/rand"
	"time"

	"github.com/cshabsin/thegrid/cardkit/card"
)

type Pile []*card.Card

func (p *Pile) Push(c *card.Card) {
	*p = append(*p, c)
}

func (p *Pile) Pop() *card.Card {
	if len(*p) == 0 {
		return nil
	}
	card := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return card
}

func (p Pile) Peek() *card.Card {
	if len(p) == 0 {
		return nil
	}
	return p[len(p)-1]
}

func (p Pile) Len() int {
	return len(p)
}

func (p Pile) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p), func(i, j int) {
		p[i], p[j] = p[j], p[i]
	})
}
