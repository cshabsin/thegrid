package ui

import (
	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/pile"
	"github.com/cshabsin/thegrid/js"
)

type LayoutDirection int

const (
	Horizontal LayoutDirection = iota
	Vertical
)

type PileLayout struct {
	Direction  LayoutDirection
	CardOffset int
	MaxVisible int
}

type Game interface {
	AllCards() []*card.Card
	GetAllPiles() map[string]pile.Pile
	GetPileLayout(name string) PileLayout
	AddListener(func())
}

type Board struct {
	game      Game
	document  js.DOMDocument
	boardDiv  js.DOMElement
	pileToDOM map[string]js.DOMElement
	cardToDOM map[*card.Card]js.DOMElement
}

func NewBoard(g Game) *Board {
	// ... implementation to follow
	return nil
}

func (b *Board) Render() {
	// ... implementation to follow
}
