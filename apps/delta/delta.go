package main

import (
	"fmt"

	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/pile"
	"github.com/cshabsin/thegrid/cardkit/ui"
	"github.com/cshabsin/thegrid/js"
)

type Game struct{}

func (g *Game) AllCards() []*card.Card                  { return nil }
func (g *Game) GetAllPiles() map[string]pile.Pile       { return nil }
func (g *Game) GetPileLayout(name string) ui.PileLayout { return ui.PileLayout{} }
func (g *Game) AddListener(func())                      {}
func (g *Game) CheckWin() bool                          { return false }
func (g *Game) SelectedCard() *card.Card                { return nil }
func (g *Game) SetSelectedCard(*card.Card)              {}
func (g *Game) MoveToFoundation(*card.Card)             {}
func (g *Game) MoveSelectedToPile(string)               {}
func (g *Game) FlipCard(*card.Card)                     {}
func (g *Game) StockClicked()                           {}
func (g *Game) ToggleDebugWin()                         {}
func (g *Game) NewGame()                                {}

func main() {
	doc := js.Document()
	doc.AddEventListener("DOMContentLoaded", func(_ js.DOMElement, _ js.DOMEvent) {
		fmt.Println("Hello, Delta!")
		game := &Game{}
		ui.NewBoard(game, doc, doc.GetElementByID("game-board"))
	})

	select {}
}
