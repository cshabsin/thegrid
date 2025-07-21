package main

import (
	"fmt"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/dragdrop"
	"github.com/cshabsin/thegrid/solitaire/game"
)

type GameUI struct {
	Board       js.DOMElement
	Stock       js.DOMElement
	Waste       js.DOMElement
	Foundations [4]js.DOMElement
	Tableau     [7]js.DOMElement
}

var g *game.Game

func createDiv(doc js.DOMDocument, attrs ...attr.Attr) js.DOMElement {
	return doc.CreateElement("div", attrs...)
}

func main() {
	g = game.New()
	document := js.Document()
	board := document.GetElementByID("game-board")
	board.Clear() // Clear the board

	ui := &GameUI{Board: board}

	// Create top row elements
	topRow := createDiv(document)
	topRow.Style().Set("grid-row", "1")
	topRow.Style().Set("grid-column", "1 / span 7")
	topRow.Style().Set("display", "grid")
	topRow.Style().Set("grid-template-columns", "repeat(7, 1fr)")
	board.Append(topRow)

	ui.Stock = createDiv(document)
	ui.Stock.Style().Set("grid-column", "1")
	ui.Stock.Style().Set("position", "relative")
	topRow.Append(ui.Stock)

	ui.Waste = createDiv(document)
	ui.Waste.Style().Set("grid-column", "2")
	ui.Waste.Style().Set("position", "relative")
	topRow.Append(ui.Waste)
	ui.Waste.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if len(g.Waste) > 0 {
			g.SelectedCard = g.Waste[len(g.Waste)-1]
			g.NotifyListeners() // This will trigger a render
		}
	})

	for i := range 4 {
		foundationIndex := i
		ui.Foundations[i] = createDiv(document)
		ui.Foundations[i].Style().Set("grid-column", fmt.Sprintf("%d", i+4))
		ui.Foundations[i].Style().Set("position", "relative")
		topRow.Append(ui.Foundations[i])
		ui.Foundations[i].AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
			if g.SelectedCard == nil {
				return
			}
			foundation := g.Foundations[foundationIndex]
			if len(foundation) == 0 {
				if g.SelectedCard.Rank == game.Ace {
					g.MoveSelectedToFoundation(foundationIndex)
				}
				return
			}
			topCard := foundation[len(foundation)-1]
			if g.SelectedCard.Suit == topCard.Suit && g.SelectedCard.Rank == topCard.Rank+1 {
				g.MoveSelectedToFoundation(foundationIndex)
			}
		})
		dragdrop.NewDropTarget(ui.Foundations[i], func(e js.DOMEvent) {
			if g.SelectedCard == nil {
				return
			}
			foundation := g.Foundations[foundationIndex]
			if len(foundation) == 0 {
				if g.SelectedCard.Rank == game.Ace {
					g.MoveSelectedToFoundation(foundationIndex)
				}
			} else {
				topCard := foundation[len(foundation)-1]
				if g.SelectedCard.Suit == topCard.Suit && g.SelectedCard.Rank == topCard.Rank+1 {
					g.MoveSelectedToFoundation(foundationIndex)
				}
			}
		})
	}

	// Create tableau elements
	tableauRow := createDiv(document)
	tableauRow.Style().Set("grid-row", "2")
	tableauRow.Style().Set("grid-column", "1 / span 7")
	tableauRow.Style().Set("display", "grid")
	tableauRow.Style().Set("grid-template-columns", "repeat(7, 1fr)")
	board.Append(tableauRow)
	for i := 0; i < 7; i++ {
		pileIndex := i
		ui.Tableau[i] = createDiv(document)
		ui.Tableau[i].Style().Set("grid-column", fmt.Sprintf("%d", i+1))
		ui.Tableau[i].Style().Set("position", "relative")
		tableauRow.Append(ui.Tableau[i])
		ui.Tableau[i].AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
			pile := g.Tableau[pileIndex]

			// First, determine which card was clicked, if any.
			var clickedCard *game.Card
			clickedCardIsLast := false

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
								g.NotifyListeners()
							}
							return
						}
						clickedCard = card
						if i == len(pile)-1 {
							clickedCardIsLast = true
						}
						break
					}
				}
			}

			if g.SelectedCard != nil {
				if clickedCard == g.SelectedCard {
					g.SelectedCard = nil
				} else {
					canMoveTo := false
					destPile := g.Tableau[pileIndex]
					if len(destPile) == 0 {
						if g.SelectedCard.Rank == game.King {
							canMoveTo = true
						}
					} else if clickedCardIsLast {
						topCard := destPile[len(destPile)-1]
						if g.SelectedCard.Suit.Color() != topCard.Suit.Color() && g.SelectedCard.Rank == topCard.Rank-1 {
							canMoveTo = true
						}
					}

					if canMoveTo {
						g.MoveSelectedToTableau(pileIndex)
					} else {
						g.SelectedCard = clickedCard
					}
				}
			} else {
				g.SelectedCard = clickedCard
			}
			g.NotifyListeners()
		})
		dragdrop.NewDropTarget(ui.Tableau[i], func(e js.DOMEvent) {
			if g.SelectedCard == nil {
				return
			}
			destPile := g.Tableau[pileIndex]
			if len(destPile) == 0 {
				if g.SelectedCard.Rank == game.King {
					g.MoveSelectedToTableau(pileIndex)
				}
			} else {
				topCard := destPile[len(destPile)-1]
				if g.SelectedCard.Suit.Color() != topCard.Suit.Color() && g.SelectedCard.Rank == topCard.Rank-1 {
					g.MoveSelectedToTableau(pileIndex)
				}
			}
		})
	}

	ui.Stock.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if len(g.Stock) > 0 {
			g.DrawCards()
		} else {
			g.RecycleWaste()
		}
	})

	g.AddListener(func() {
		render(document, ui, g)
	})
	render(document, ui, g)
	select {}
}

func render(document js.DOMDocument, ui *GameUI, g *game.Game) {
	if g.CheckWin() {
		winDiv := createDiv(document)
		winDiv.Style().Set("position", "absolute")
		winDiv.Style().Set("top", "50%")
		winDiv.Style().Set("left", "50%")
		winDiv.Style().Set("transform", "translate(-50%, -50%)")
		winDiv.Style().Set("font-size", "5em")
		winDiv.Style().Set("color", "white")
		winDiv.Append(document.CreateElement("h1").SetText("You Win!"))
		ui.Board.Append(winDiv)
		return
	}
	// Render Stock
	ui.Stock.Clear()
	stockCardDiv := createDiv(document, attr.Class("card"))
	if len(g.Stock) > 0 {
		stockCardDiv.Style().Set("background-color", "blue")
	} else {
		stockCardDiv.Style().Set("border", "1px solid black")
	}
	ui.Stock.Append(stockCardDiv)

	// Render Waste
	ui.Waste.Clear()
	wastePlaceholder := createDiv(document, attr.Class("card"))
	wastePlaceholder.Style().Set("border", "1px solid black")
	ui.Waste.Append(wastePlaceholder)
	if len(g.Waste) > 0 {
		start := len(g.Waste) - 3
		if start < 0 {
			start = 0
		}
		for i := start; i < len(g.Waste); i++ {
			card := g.Waste[i]
			cardDiv := createDiv(document, attr.Class("card"))
			if i == len(g.Waste)-1 && card == g.SelectedCard {
				cardDiv.Style().Set("border", "2px solid yellow")
			}
			cardDiv.Style().Set("left", fmt.Sprintf("%dpx", (i-start)*20))
			cardDiv.SetAttr("draggable", true)
			suitDiv := createDiv(document, attr.Class("suit"))
			suitDiv.Style().Set("color", card.Suit.Color())
			suitDiv.SetText(card.Suit.String())
			cardDiv.Append(suitDiv)
			rankDiv := createDiv(document, attr.Class("rank"))
			rankDiv.Style().Set("color", card.Suit.Color())
			rankDiv.SetText(card.Rank.String())
			cardDiv.Append(rankDiv)
			ui.Waste.Append(cardDiv)
			cardDiv.AddEventListener("dblclick", func(el js.DOMElement, e js.DOMEvent) {
				for i := 0; i < 4; i++ {
					foundation := g.Foundations[i]
					if len(foundation) == 0 {
						if card.Rank == game.Ace {
							g.SelectedCard = card
							g.MoveSelectedToFoundation(i)
							return
						}
					} else {
						topCard := foundation[len(foundation)-1]
						if card.Suit == topCard.Suit && card.Rank == topCard.Rank+1 {
							g.SelectedCard = card
							g.MoveSelectedToFoundation(i)
							return
						}
					}
				}
			})
			dragdrop.NewDraggable(cardDiv, func(e js.DOMEvent) {
				g.SelectedCard = card
			})
		}
	}

	// Render Foundations
	for i := 0; i < 4; i++ {
		foundationDiv := ui.Foundations[i]
		foundationDiv.Clear()
		placeholder := createDiv(document, attr.Class("card"))
		placeholder.Style().Set("border", "1px solid black")
		foundationDiv.Append(placeholder)
		if len(g.Foundations[i]) > 0 {
			card := g.Foundations[i][len(g.Foundations[i])-1]
			cardDiv := createDiv(document, attr.Class("card"))
			suitDiv := createDiv(document, attr.Class("suit"))
			suitDiv.Style().Set("color", card.Suit.Color())
			suitDiv.SetText(card.Suit.String())
			cardDiv.Append(suitDiv)
			rankDiv := createDiv(document, attr.Class("rank"))
			rankDiv.Style().Set("color", card.Suit.Color())
			rankDiv.SetText(card.Rank.String())
			cardDiv.Append(rankDiv)
			foundationDiv.Append(cardDiv)
		}
	}

	// Render Tableau
	for i, pile := range g.Tableau {
		pileDiv := ui.Tableau[i]
		pileDiv.Clear()
		for j, card := range pile {
			cardDiv := createDiv(document, attr.Class("card"))
			if card == g.SelectedCard {
				cardDiv.Style().Set("border", "2px solid yellow")
			}
			cardDiv.Style().Set("top", fmt.Sprintf("%dpx", j*30))
			if card.FaceUp {
				cardDiv.SetAttr("draggable", true)
				suitDiv := createDiv(document, attr.Class("suit"))
				suitDiv.Style().Set("color", card.Suit.Color())
				suitDiv.SetText(card.Suit.String())
				cardDiv.Append(suitDiv)
				rankDiv := createDiv(document, attr.Class("rank"))
				rankDiv.Style().Set("color", card.Suit.Color())
				rankDiv.SetText(card.Rank.String())
				cardDiv.Append(rankDiv)
			} else {
				cardDiv.Style().Set("background-color", "blue")
			}
			pileDiv.Append(cardDiv)
			cardDiv.AddEventListener("dblclick", func(el js.DOMElement, e js.DOMEvent) {
				if !card.FaceUp {
					return
				}
				for i := 0; i < 4; i++ {
					foundation := g.Foundations[i]
					if len(foundation) == 0 {
						if card.Rank == game.Ace {
							g.SelectedCard = card
							g.MoveSelectedToFoundation(i)
							return
						}
					} else {
						topCard := foundation[len(foundation)-1]
						if card.Suit == topCard.Suit && card.Rank == topCard.Rank+1 {
							g.SelectedCard = card
							g.MoveSelectedToFoundation(i)
							return
						}
					}
				}
			})
			if card.FaceUp {
				dragdrop.NewDraggable(cardDiv, func(e js.DOMEvent) {
					g.SelectedCard = card
				})
			}
		}
	}
}
