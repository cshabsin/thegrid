package main

import (
	"fmt"
	"math"

	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/deck"
	"github.com/cshabsin/thegrid/cardkit/pile"
	"github.com/cshabsin/thegrid/cardkit/ui"
	"github.com/cshabsin/thegrid/js"
)

const numPlayers = 6

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

	for i := range numPlayers {
		for range 2 {
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
		CardOffset: 25,
	}
	if player == g.povPlayer {
		layout.CardOffset = 25
		layout.Top = 600
		layout.Left = 495
	} else {
		// Arrange other players in a semicircle at the top of the board.
		numOtherPlayers := numPlayers - 1
		playerIndex := player
		if player > g.povPlayer {
			playerIndex--
		}
		angle := math.Pi * float64(playerIndex+1) / float64(numOtherPlayers+1)
		radius := 300.0
		layout.Left = int(512 - radius*math.Cos(angle) - 50)
		layout.Top = int(325 - radius*math.Sin(angle))
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
	setup := func() {
		fmt.Println("Hello, Delta!")
		game := New(0)
		doc := js.Document()
		boardDiv := doc.GetElementByID("game-board")
		ui.NewBoard(game, doc, boardDiv)

		loginButton := doc.GetElementByID("login-button")
		loginButton.AddEventListener("click", func(_ js.DOMElement, _ js.DOMEvent) {
			js.Global().Call("signIn")
		})

		js.Global().Call("onAuthStateChanged", js.FuncOf(func(this js.Value, args []js.Value) any {
			user := args[0]
			if !user.IsNull() {
				fmt.Println("User is signed in:", user.Get("displayName").String())
			} else {
				fmt.Println("User is signed out")
			}
			return nil
		}))
	}

	doc := js.Document()
	if doc.ReadyState() == "loading" {
		doc.AddEventListener("DOMContentLoaded", func(_ js.DOMElement, _ js.DOMEvent) {
			setup()
		})
	} else {
		setup()
	}

	select {}
}
