package klondike

import (
	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/pile"
)

type Klondike struct {
	Deck         card.Deck
	Tableau      [7]pile.Pile
	Foundations  [4]pile.Pile
	Stock        pile.Pile
	Waste        pile.Pile
	SelectedCard *card.Card
	listeners    []func()
}

func New() *Klondike {
	deck := card.NewStandard52()
	deck.Shuffle()

	game := &Klondike{
		Deck: deck,
	}

	// Deal to tableau
	for i := 0; i < 7; i++ { // i is the pile index
		for j := 0; j <= i; j++ { // j is the card index within the pile
			card := game.Deck.Draw()
			if j == i { // Only the last card in the pile is face up
				card.FaceUp = true
			}
			game.Tableau[i].Push(card)
		}
	}

	game.Stock = pile.Pile(game.Deck)
	game.Deck = nil

	return game
}

func (g *Klondike) RecycleWaste() {
	if g.Stock.Len() > 0 {
		return
	}
	g.Stock = g.Waste
	g.Waste = nil
	for _, card := range g.Stock {
		card.FaceUp = false
	}
	g.NotifyListeners()
}

func (g *Klondike) DrawCards() {
	numToDraw := 3
	if g.Stock.Len() < 3 {
		numToDraw = g.Stock.Len()
	}
	for i := 0; i < numToDraw; i++ {
		card := g.Stock.Pop()
		card.FaceUp = true
		g.Waste.Push(card)
	}
	g.NotifyListeners()
}

// findAndRemoveSelectedCard finds the selected card, removes it (and any cards on top of it) from its current location,
// and returns the stack of cards that was removed. Returns nil if no card was selected or found.
func (g *Klondike) findAndRemoveSelectedCard() pile.Pile {
	if g.SelectedCard == nil {
		return nil
	}

	// Find in tableau
	for i, p := range g.Tableau {
		for j, c := range p {
			if c == g.SelectedCard {
				stack := g.Tableau[i][j:]
				g.Tableau[i] = g.Tableau[i][:j]
				return stack
			}
		}
	}

	// Find in waste
	if g.Waste.Len() > 0 && g.Waste.Peek() == g.SelectedCard {
		return pile.Pile{g.Waste.Pop()}
	}

	return nil
}

func (g *Klondike) MoveSelectedToFoundation(foundationIndex int) {
	if g.SelectedCard == nil {
		return
	}
	foundation := &g.Foundations[foundationIndex]
	// Foundation move is only valid if moving a single card.
	// Check if the selected card is the top of its stack.
	if g.Waste.Len() > 0 && g.Waste.Peek() == g.SelectedCard {
		// It's from the waste pile, which is a single card move.
	} else {
		found := false
		for _, p := range g.Tableau {
			if p.Len() > 0 && p.Peek() == g.SelectedCard {
				found = true
				break
			}
		}
		if !found {
			// The selected card is not at the top of a tableau pile, so it's part of a stack. Invalid move.
			return
		}
	}

	if foundation.Len() == 0 {
		if g.SelectedCard.Rank != card.Ace {
			return // Must be an Ace on an empty foundation
		}
	} else {
		topCard := foundation.Peek()
		if g.SelectedCard.Suit != topCard.Suit || g.SelectedCard.Rank != topCard.Rank+1 {
			return // Must be same suit and next rank
		}
	}

	stack := g.findAndRemoveSelectedCard()
	if stack != nil {
		*foundation = append(*foundation, stack...)
		g.SelectedCard = nil
		g.NotifyListeners()
	}
}

func (g *Klondike) MoveSelectedToTableau(tableauIndex int) {
	if g.SelectedCard == nil {
		return
	}
	destPile := &g.Tableau[tableauIndex]
	if destPile.Len() > 0 {
		topCard := destPile.Peek()
		if g.SelectedCard.Suit.Color() == topCard.Suit.Color() || g.SelectedCard.Rank != topCard.Rank-1 {
			return // Invalid move
		}
	} else {
		if g.SelectedCard.Rank != card.King {
			return // Must be a king on an empty pile
		}
	}

	stack := g.findAndRemoveSelectedCard()
	if stack != nil {
		*destPile = append(*destPile, stack...)
		g.SelectedCard = nil
		g.NotifyListeners()
	}
}

func (g *Klondike) CheckWin() bool {
	for _, f := range g.Foundations {
		if f.Len() != 13 {
			return false
		}
	}
	return true
}

func (g *Klondike) AddListener(listener func()) {
	g.listeners = append(g.listeners, listener)
}

func (g *Klondike) NotifyListeners() {
	for _, listener := range g.listeners {
		listener()
	}
}