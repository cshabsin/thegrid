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
}

var klondikeGame *klondike.Klondike

func createDiv(doc js.DOMDocument, attrs ...attr.Attr) js.DOMElement {
	return doc.CreateElement("div", attrs...)
}

func main() {
	klondikeGame = klondike.New()
	document := js.Document()
	board := document.GetElementByID("game-board")
	board.Clear() // Clear the board

	ui := &GameUI{Board: board}

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
	for i := 0; i < 7; i++ {
		pileIndex := i
		ui.Tableau[i] = createDiv(document, attr.Class("pile")).SetStyle(style.GridColumn(fmt.Sprintf("%d", i+1)))
		tableauRow.Append(ui.Tableau[i])
		ui.Tableau[i].AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
			pile := klondikeGame.Tableau[pileIndex]

			// First, determine which card was clicked, if any.
			var clickedCard *card.Card

			if len(pile) > 0 {
				clientY := e.Value.Get("clientY").Int()
				cardElements := el.QuerySelectorAll(".card")
				for i := len(pile) - 1; i >= 0; i-- {
					card := pile[i]
					cardEl := cardElements[i]
					rect := cardEl.GetBoundingClientRect()
					if clientY >= rect.Get("top").Int() && clientY <= rect.Get("bottom").Int() {
						if !card.FaceUp {
							if i == len(pile)-1 {
								card.FaceUp = true
								klondikeGame.NotifyListeners()
							}
							return
						}
						clickedCard = card
						break
					}
				}
			}

			if klondikeGame.SelectedCard != nil {
				if clickedCard == klondikeGame.SelectedCard {
					klondikeGame.SelectedCard = nil
				} else {
					if klondikeGame.CanMoveToTableau(klondikeGame.SelectedCard, pileIndex) {
						klondikeGame.MoveSelectedToTableau(pileIndex)
					} else {
						klondikeGame.SelectedCard = clickedCard
					}
				}
			} else {
				klondikeGame.SelectedCard = clickedCard
			}
			klondikeGame.NotifyListeners()
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
	if klondikeGame.CheckWin() {
		winDiv := createDiv(document, attr.Class("win-div"))
		winDiv.Append(document.CreateElement("h1").SetText("You Win!"))
		ui.Board.Append(winDiv)
		return
	}
	// Render Stock
	ui.Stock.Clear()
	stockCardDiv := createDiv(document, attr.Class("card"))
	if len(klondikeGame.Stock) > 0 {
		stockCardDiv.AddClass("face-down-card")
	} else {
		stockCardDiv.AddClass("card-placeholder")
	}
	ui.Stock.Append(stockCardDiv)

	// Render Waste
	ui.Waste.Clear()
	if len(klondikeGame.Waste) == 0 {
		wastePlaceholder := createDiv(document, attr.Class("card-placeholder"))
		ui.Waste.Append(wastePlaceholder)
	} else {
		start := len(klondikeGame.Waste) - 3
		if start < 0 {
			start = 0
		}
		for i := start; i < len(klondikeGame.Waste); i++ {
			wasteCard := klondikeGame.Waste[i]
			cardDiv := createDiv(document, attr.Class("card"))
			cardDiv.AddClass("face-up-card")
			if i == len(klondikeGame.Waste)-1 && wasteCard == klondikeGame.SelectedCard {
				cardDiv.AddClass("selected-card")
			}
			cardDiv.SetStyle(style.Left(fmt.Sprintf("%dpx", (i-start)*20)))
			cardDiv.SetAttr("draggable", true)
			suitDiv := createDiv(document, attr.Class("suit")).SetStyle(style.Color(wasteCard.Suit.Color()))
			suitDiv.SetText(wasteCard.Suit.String())
			cardDiv.Append(suitDiv)
			rankDiv := createDiv(document, attr.Class("rank")).SetStyle(style.Color(wasteCard.Suit.Color()))
			rankDiv.SetText(wasteCard.Rank.String())
			cardDiv.Append(rankDiv)
			ui.Waste.Append(cardDiv)
			cardDiv.AddEventListener("dblclick", func(el js.DOMElement, e js.DOMEvent) {
				for i := range klondikeGame.Foundations {
					if klondikeGame.CanMoveToFoundation(wasteCard, i) {
						klondikeGame.SelectedCard = wasteCard
						klondikeGame.MoveSelectedToFoundation(i)
						return
					}
				}
			})
			dragdrop.NewDraggable(cardDiv, func(e js.DOMEvent) {
				klondikeGame.SelectedCard = wasteCard
			})
		}
	}

	// Render Foundations
	for i := range ui.Foundations {
		foundationDiv := ui.Foundations[i]
		foundationDiv.Clear()
		if len(klondikeGame.Foundations[i]) == 0 {
			placeholder := createDiv(document, attr.Class("card-placeholder"))
			foundationDiv.Append(placeholder)
		} else {
			card := klondikeGame.Foundations[i][len(klondikeGame.Foundations[i])-1]
			cardDiv := createDiv(document, attr.Class("card"))
			cardDiv.AddClass("face-up-card")
			suitDiv := createDiv(document, attr.Class("suit")).SetStyle(style.Color(card.Suit.Color()))
			suitDiv.SetText(card.Suit.String())
			cardDiv.Append(suitDiv)
			rankDiv := createDiv(document, attr.Class("rank")).SetStyle(style.Color(card.Suit.Color()))
			rankDiv.SetText(card.Rank.String())
			cardDiv.Append(rankDiv)
			foundationDiv.Append(cardDiv)
		}
	}

	// Render Tableau
	for i, pile := range klondikeGame.Tableau {
		pileDiv := ui.Tableau[i]
		pileDiv.Clear()
		if len(pile) == 0 {
			placeholder := createDiv(document, attr.Class("card-placeholder"))
			pileDiv.Append(placeholder)
		}
		for j, currentCard := range pile {
			cardDiv := createDiv(document, attr.Class("card"))
			if currentCard == klondikeGame.SelectedCard {
				cardDiv.AddClass("selected-card")
			}
			cardDiv.SetStyle(style.Top(fmt.Sprintf("%dpx", j*30)))
			if currentCard.FaceUp {
				cardDiv.AddClass("face-up-card")
				cardDiv.SetAttr("draggable", true)
				suitDiv := createDiv(document, attr.Class("suit")).SetStyle(style.Color(currentCard.Suit.Color()))
				suitDiv.SetText(currentCard.Suit.String())
				cardDiv.Append(suitDiv)
				rankDiv := createDiv(document, attr.Class("rank")).SetStyle(style.Color(currentCard.Suit.Color()))
				rankDiv.SetText(currentCard.Rank.String())
				cardDiv.Append(rankDiv)
			} else {
				cardDiv.AddClass("face-down-card")
			}
			pileDiv.Append(cardDiv)
			cardDiv.AddEventListener("dblclick", func(el js.DOMElement, e js.DOMEvent) {
				if !currentCard.FaceUp {
					return
				}
				for i := range klondikeGame.Foundations {
					if klondikeGame.CanMoveToFoundation(currentCard, i) {
						klondikeGame.SelectedCard = currentCard
						klondikeGame.MoveSelectedToFoundation(i)
						return
					}
				}
			})
			if currentCard.FaceUp {
				dragdrop.NewDraggable(cardDiv, func(e js.DOMEvent) {
					klondikeGame.SelectedCard = currentCard
				})
			}
		}
	}
}
