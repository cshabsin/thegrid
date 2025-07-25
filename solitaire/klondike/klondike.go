package klondike

import (
	"fmt"
	"strings"

	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/pile"
	"github.com/cshabsin/thegrid/cardkit/ui"
)

type Klondike struct {
	Deck         card.Deck
	Tableau      [7]pile.Pile
	Foundations  [4]pile.Pile
	Stock        pile.Pile
	Waste        pile.Pile
	selectedCard *card.Card
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
	if g.selectedCard == nil {
		return nil
	}

	// Find in tableau
	for i, p := range g.Tableau {
		for j, c := range p {
			if c == g.selectedCard {
				stack := g.Tableau[i][j:]
				g.Tableau[i] = g.Tableau[i][:j]
				return stack
			}
		}
	}

	// Find in waste
	if g.Waste.Len() > 0 && g.Waste.Peek() == g.selectedCard {
		return pile.Pile{g.Waste.Pop()}
	}

	return nil
}

func (g *Klondike) MoveSelectedToFoundation(foundationIndex int) {
	if !g.CanMoveToFoundation(g.selectedCard, foundationIndex) {
		return
	}

	stack := g.findAndRemoveSelectedCard()
	if stack != nil {
		g.Foundations[foundationIndex].Push(stack[0])
		g.selectedCard = nil
		g.NotifyListeners()
	}
}

func (g *Klondike) CanMoveToFoundation(c *card.Card, foundationIndex int) bool {
	if c == nil {
		return false
	}
	foundation := &g.Foundations[foundationIndex]
	if foundation.Len() == 0 {
		return c.Rank == card.Ace
	}
	topCard := foundation.Peek()
	return c.Suit == topCard.Suit && c.Rank == topCard.Rank+1
}

func (g *Klondike) MoveSelectedToTableau(tableauIndex int) {
	if !g.CanMoveToTableau(g.selectedCard, tableauIndex) {
		return
	}

	stack := g.findAndRemoveSelectedCard()
	if stack != nil {
		g.Tableau[tableauIndex] = append(g.Tableau[tableauIndex], stack...)
		g.selectedCard = nil
		g.NotifyListeners()
	}
}

func (g *Klondike) CanMoveToTableau(c *card.Card, tableauIndex int) bool {
	if c == nil {
		return false
	}
	destPile := &g.Tableau[tableauIndex]
	if destPile.Len() == 0 {
		return c.Rank == card.King
	}
	topCard := destPile.Peek()
	return c.Suit.Color() != topCard.Suit.Color() && c.Rank == topCard.Rank-1
}

func (g *Klondike) AllCards() []*card.Card {
	var cards []*card.Card
	for _, p := range g.Tableau {
		cards = append(cards, p...)
	}
	for _, p := range g.Foundations {
		cards = append(cards, p...)
	}
	cards = append(cards, g.Stock...)
	cards = append(cards, g.Waste...)
	return cards
}

func (g *Klondike) SelectedCard() *card.Card {
	return g.selectedCard
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

func (g *Klondike) GetAllPiles() map[string]pile.Pile {
	piles := make(map[string]pile.Pile)
	piles["stock"] = g.Stock
	piles["waste"] = g.Waste
	for i, p := range g.Foundations {
		piles[fmt.Sprintf("foundation-%d", i)] = p
	}
	for i, p := range g.Tableau {
		piles[fmt.Sprintf("tableau-%d", i)] = p
	}
	return piles
}

func (g *Klondike) GetPileLayout(name string) ui.PileLayout {
	switch name {
	case "stock":
		return ui.PileLayout{GridRow: 1, GridColumn: 1}
	case "waste":
		return ui.PileLayout{GridRow: 1, GridColumn: 2, Direction: ui.Horizontal, CardOffset: 20, MaxVisible: 3}
	default:
		if strings.HasPrefix(name, "foundation-") {
			var i int
			fmt.Sscanf(name, "foundation-%d", &i)
			return ui.PileLayout{GridRow: 1, GridColumn: i + 4}
		}
		if strings.HasPrefix(name, "tableau-") {
			var i int
			fmt.Sscanf(name, "tableau-%d", &i)
			return ui.PileLayout{GridRow: 2, GridColumn: i + 1, Direction: ui.Vertical, CardOffset: 30}
		}
		return ui.PileLayout{}
	}
}