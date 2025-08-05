package main

import (
	"fmt"

	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/deck"
	"github.com/cshabsin/thegrid/cardkit/pile"
	"github.com/cshabsin/thegrid/cardkit/ui"
	"github.com/cshabsin/thegrid/js"
)

const numPlayers = 2

type Game struct {
	deck      deck.Deck
	hands     [numPlayers]pile.Pile
	povPlayer int
	listeners []func()
}

func New(povPlayer int) *Game {
	deck := deck.NewStandard52()
	deck.Shuffle()

	game := &Game{
		deck:      deck,
		povPlayer: povPlayer,
	}

	for i := 0; i < numPlayers; i++ {
		for j := 0; j < 2; j++ {
			card := game.deck.Draw()
			if i == game.povPlayer {
				card.FaceUp = true
			}
			game.hands[i].Push(card)
		}
	}

	return game
}

func (g *Game) AllCards() []*card.Card {
	var cards []*card.Card
	for _, hand := range g.hands {
		cards = append(cards, hand...)
	}
	cards = append(cards, g.deck...)
	return cards
}

func (g *Game) GetAllPiles() map[string]pile.Pile {
	piles := make(map[string]pile.Pile)
	for i, hand := range g.hands {
		piles[fmt.Sprintf("hand-%d", i)] = hand
	}
	return piles
}

func (g *Game) GetPileLayout(name string) ui.PileLayout {
	var player int
	fmt.Sscanf(name, "hand-%d", &player)
	layout := ui.PileLayout{
		Direction:  ui.Horizontal,
		CardOffset: 20,
	}
	if player == g.povPlayer {
		layout.CardOffset = 110
		layout.ClassName = fmt.Sprintf("hand-%d", player)
	} else {
		layout.GridRow = 1
		layout.GridColumn = player + 1
	}
	return layout
}

func (g *Game) AddListener(listener func()) {
	g.listeners = append(g.listeners, listener)
}

func (g *Game) NotifyListeners() {
	for _, listener := range g.listeners {
		listener()
	}
}

func (g *Game) CheckWin() bool              { return false }
func (g *Game) SelectedCard() *card.Card    { return nil }
func (g *Game) SetSelectedCard(*card.Card)  {}
func (g *Game) MoveToFoundation(*card.Card) {}
func (g *Game) MoveSelectedToPile(string)   {}
func (g *Game) FlipCard(*card.Card)         {}
func (g *Game) StockClicked()               {}
func (g *Game) ToggleDebugWin()             {}
func (g *Game) NewGame()                    {}

func main() {
	doc := js.Document()
	doc.AddEventListener("DOMContentLoaded", func(_ js.DOMElement, _ js.DOMEvent) {
		fmt.Println("Hello, Delta!")
		game := New(0)
		ui.NewBoard(game, doc, doc.GetElementByID("game-board"))
	})

	select {}
}
