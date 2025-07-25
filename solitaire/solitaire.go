package main

import (
	"fmt"

	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/dragdrop"
	"github.com/cshabsin/thegrid/js/style"
	"github.com/cshabsin/thegrid/solitaire/klondike"
)

type GameUI struct {
	Board       js.DOMElement
	Stock       js.DOMElement
	Waste       js.DOMElement
	Foundations [4]js.DOMElement
	Tableau     [7]js.DOMElement
	CardToDOM   map[*card.Card]js.DOMElement
}

var klondikeGame *klondike.Klondike

func createCardElement(doc js.DOMDocument, c *card.Card) js.DOMElement {
	cardDiv := createDiv(doc, attr.Class("card"))
	dragdrop.NewDraggable(cardDiv, func(e js.DOMEvent) {
		klondikeGame.SelectedCard = c
	})
	cardDiv.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if klondikeGame.SelectedCard == c {
			klondikeGame.SelectedCard = nil
		} else {
			klondikeGame.SelectedCard = c
		}
		klondikeGame.NotifyListeners()
	})
	cardDiv.AddEventListener("dblclick", func(el js.DOMElement, e js.DOMEvent) {
		if !c.FaceUp {
			return
		}
		for i := range klondikeGame.Foundations {
			if klondikeGame.CanMoveToFoundation(c, i) {
				klondikeGame.SelectedCard = c
				klondikeGame.MoveSelectedToFoundation(i)
				return
			}
		}
	})
	return cardDiv
}

func createDiv(doc js.DOMDocument, attrs ...attr.Attr) js.DOMElement {
	return doc.CreateElement("div", attrs...)
}

func populateCardElement(doc js.DOMDocument, cardDiv js.DOMElement, c *card.Card) {
	suitDiv := createDiv(doc, attr.Class("suit")).SetStyle(style.Color(c.Suit.Color()))
	suitDiv.SetText(c.Suit.String())
	cardDiv.Append(suitDiv)
	rankDiv := createDiv(doc, attr.Class("rank")).SetStyle(style.Color(c.Suit.Color()))
	rankDiv.SetText(c.Rank.String())
	cardDiv.Append(rankDiv)
}

func main() {
	klondikeGame = klondike.New()
	document := js.Document()
	board := document.GetElementByID("game-board")
	board.Clear() // Clear the board

	ui := &GameUI{
		Board:     board,
		CardToDOM: make(map[*card.Card]js.DOMElement),
	}

	// Create all card elements upfront
	for _, c := range klondikeGame.AllCards() {
		ui.CardToDOM[c] = createCardElement(document, c)
	}

	// Create top row elements
	topRow := createDiv(document, attr.Class("top-row"))
	board.Append(topRow)

	ui.Stock = createDiv(document, attr.Class("pile")).SetStyle(style.GridColumn("1"))
	topRow.Append(ui.Stock)

	ui.Waste = createDiv(document, attr.Class("pile")).SetStyle(style.GridColumn("2"))
	topRow.Append(ui.Waste)
	ui.Waste.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if len(klondikeGame.Waste) > 0 {
			klondikeGame.SelectedCard = klondikeGame.Waste[len(klondikeGame.Waste)-1]
			klondikeGame.NotifyListeners() // This will trigger a render
		}
	})

	for i := range ui.Foundations {
		foundationIndex := i
		ui.Foundations[i] = createDiv(document, attr.Class("pile")).SetStyle(style.GridColumn(fmt.Sprintf("%d", i+4)))
		topRow.Append(ui.Foundations[i])
		ui.Foundations[i].AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
			if klondikeGame.CanMoveToFoundation(klondikeGame.SelectedCard, foundationIndex) {
				klondikeGame.MoveSelectedToFoundation(foundationIndex)
			}
		})
		foundationDropTarget := dragdrop.NewDropTarget(ui.Foundations[i], func(e js.DOMEvent) {
			if klondikeGame.CanMoveToFoundation(klondikeGame.SelectedCard, foundationIndex) {
				klondikeGame.MoveSelectedToFoundation(foundationIndex)
			}
		})
		foundationDropTarget.CanDrop = func(e js.DOMEvent) bool {
			return klondikeGame.CanMoveToFoundation(klondikeGame.SelectedCard, foundationIndex)
		}
	}

	// Create tableau elements
	tableauRow := createDiv(document, attr.Class("tableau-row"))
	board.Append(tableauRow)
	for i := range ui.Tableau {
		pileIndex := i
		ui.Tableau[i] = createDiv(document, attr.Class("pile")).SetStyle(style.GridColumn(fmt.Sprintf("%d", i+1)))
		tableauRow.Append(ui.Tableau[i])
		ui.Tableau[i].AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
			pile := klondikeGame.Tableau[pileIndex]
			if pile.Len() > 0 {
				card := pile.Peek()
				if !card.FaceUp {
					card.FaceUp = true
					klondikeGame.NotifyListeners()
				}
			}
		})
		tableauDropTarget := dragdrop.NewDropTarget(ui.Tableau[i], func(e js.DOMEvent) {
			if klondikeGame.CanMoveToTableau(klondikeGame.SelectedCard, pileIndex) {
				klondikeGame.MoveSelectedToTableau(pileIndex)
			}
		})
		tableauDropTarget.CanDrop = func(e js.DOMEvent) bool {
			return klondikeGame.CanMoveToTableau(klondikeGame.SelectedCard, pileIndex)
		}
	}

	ui.Stock.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if len(klondikeGame.Stock) > 0 {
			klondikeGame.DrawCards()
		} else {
			klondikeGame.RecycleWaste()
		}
	})

	klondikeGame.AddListener(func() {
		render(document, ui, klondikeGame)
	})
	render(document, ui, klondikeGame)
	select {}
}

func render(document js.DOMDocument, ui *GameUI, g *klondike.Klondike) {
	if g.CheckWin() {
		winDiv := createDiv(document, attr.Class("win-div"))
		winDiv.Append(document.CreateElement("h1").SetText("You Win!"))
		ui.Board.Append(winDiv)
		return
	}

	for _, c := range g.AllCards() {
		cardDiv := ui.CardToDOM[c]
		if c == g.SelectedCard {
			cardDiv.AddClass("selected-card")
		} else {
			cardDiv.RemoveClass("selected-card")
		}
	}

	// Render Stock
	ui.Stock.Clear()
	if len(g.Stock) > 0 {
		stockCardDiv := ui.CardToDOM[g.Stock[len(g.Stock)-1]]
		stockCardDiv.Clear()
		stockCardDiv.RemoveClass("face-up-card")
		stockCardDiv.AddClass("face-down-card")
		ui.Stock.Append(stockCardDiv)
	} else {
		placeholder := createDiv(document, attr.Class("card-placeholder"))
		ui.Stock.Append(placeholder)
	}

	// Render Waste
	ui.Waste.Clear()
	if len(g.Waste) > 0 {
		start := len(g.Waste) - 3
		if start < 0 {
			start = 0
		}
		for i := start; i < len(g.Waste); i++ {
			card := g.Waste[i]
			cardDiv := ui.CardToDOM[card]
			cardDiv.Clear()
			cardDiv.RemoveClass("face-down-card")
			cardDiv.AddClass("face-up-card")

			cardDiv.SetStyle(style.Left(fmt.Sprintf("%dpx", (i-start)*20)))
			cardDiv.SetAttr("draggable", true)
			populateCardElement(document, cardDiv, card)
			ui.Waste.Append(cardDiv)
		}
	} else {
		placeholder := createDiv(document, attr.Class("card-placeholder"))
		ui.Waste.Append(placeholder)
	}

	// Render Foundations
	for i := range ui.Foundations {
		foundationDiv := ui.Foundations[i]
		foundationDiv.Clear()
		if len(g.Foundations[i]) == 0 {
			placeholder := createDiv(document, attr.Class("card-placeholder"))
			foundationDiv.Append(placeholder)
		} else {
			card := g.Foundations[i].Peek()
			cardDiv := ui.CardToDOM[card]
			cardDiv.Clear()
			cardDiv.RemoveClass("face-down-card")
			cardDiv.AddClass("face-up-card")
			populateCardElement(document, cardDiv, card)
			foundationDiv.Append(cardDiv)
		}
	}

	// Render Tableau
	for i, pile := range g.Tableau {
		pileDiv := ui.Tableau[i]
		pileDiv.Clear()
		if len(pile) == 0 {
			placeholder := createDiv(document, attr.Class("card-placeholder"))
			pileDiv.Append(placeholder)
		}
		for j, card := range pile {
			cardDiv := ui.CardToDOM[card]
			cardDiv.Clear()

			cardDiv.SetStyle(style.Top(fmt.Sprintf("%dpx", j*30)), style.Left("0"))
			if card.FaceUp {
				cardDiv.SetAttr("draggable", true)
				cardDiv.RemoveClass("face-down-card")
				cardDiv.AddClass("face-up-card")
				populateCardElement(document, cardDiv, card)
			} else {
				cardDiv.SetAttr("draggable", false)
				cardDiv.RemoveClass("face-up-card")
				cardDiv.AddClass("face-down-card")
			}
			pileDiv.Append(cardDiv)
		}
	}
}
