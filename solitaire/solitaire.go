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
	board.Value.Set("innerHTML", "") // Clear the board

	ui := &GameUI{Board: board}

	// Create top row elements
	topRow := createDiv(document, attr.Style("grid-row: 1; grid-column: 1 / span 7; display: grid; grid-template-columns: repeat(7, 1fr);"))
	board.Append(topRow)

	ui.Stock = createDiv(document, attr.Style("grid-column: 1; position: relative;"))
	topRow.Append(ui.Stock)

	ui.Waste = createDiv(document, attr.Style("grid-column: 2; position: relative;"))
	topRow.Append(ui.Waste)
	ui.Waste.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if len(g.Waste) > 0 {
			g.SelectedCard = g.Waste[len(g.Waste)-1]
			g.NotifyListeners() // This will trigger a render
		}
	})

	for i := range 4 {
		foundationIndex := i
		ui.Foundations[i] = createDiv(document, attr.Style(fmt.Sprintf("grid-column: %d; position: relative;", i+4)))
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
	tableauRow := createDiv(document, attr.Style("grid-row: 2; grid-column: 1 / span 7; display: grid; grid-template-columns: repeat(7, 1fr);"))
	board.Append(tableauRow)
	for i := 0; i < 7; i++ {
		pileIndex := i
		ui.Tableau[i] = createDiv(document, attr.Style(fmt.Sprintf("grid-column: %d; position: relative;", i+1)))
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
		winDiv := createDiv(document, attr.Style("position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); font-size: 5em; color: white;"))
		winDiv.Value.Set("innerHTML", "<h1>You Win!</h1>")
		ui.Board.Append(winDiv)
		return
	}
	// Render Stock
	ui.Stock.Value.Set("innerHTML", "")
	var stockAttrs []attr.Attr
	stockAttrs = append(stockAttrs, attr.Class("card"))
	if len(g.Stock) > 0 {
		stockAttrs = append(stockAttrs, attr.Style("background-color: blue;"))
	} else {
		stockAttrs = append(stockAttrs, attr.Style("border: 1px solid black;"))
	}
	stockCardDiv := createDiv(document, stockAttrs...)
	ui.Stock.Append(stockCardDiv)

	// Render Waste
	ui.Waste.Value.Set("innerHTML", "")
	wastePlaceholder := createDiv(document, attr.Class("card"), attr.Style("border: 1px solid black;"))
	ui.Waste.Append(wastePlaceholder)
	if len(g.Waste) > 0 {
		start := len(g.Waste) - 3
		if start < 0 {
			start = 0
		}
		for i := start; i < len(g.Waste); i++ {
			card := g.Waste[i]
			var cardAttrs []attr.Attr
			cardAttrs = append(cardAttrs, attr.Class("card"))
			if i == len(g.Waste)-1 && card == g.SelectedCard {
				cardAttrs = append(cardAttrs, attr.Style("border: 2px solid yellow;"))
			}
			cardAttrs = append(cardAttrs, attr.Style(fmt.Sprintf("left: %dpx;", (i-start)*20)))
			cardAttrs = append(cardAttrs, attr.Draggable(true))
			cardDiv := createDiv(document, cardAttrs...)
			cardDiv.Value.Set("innerHTML", fmt.Sprintf(`<div class="suit" style="color:%s">%s</div><div class="rank" style="color:%s">%s</div>`, card.Suit.Color(), card.Suit.String(), card.Suit.Color(), card.Rank.String()))
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
		foundationDiv.Value.Set("innerHTML", "")
		placeholder := createDiv(document, attr.Class("card"), attr.Style("border: 1px solid black;"))
		foundationDiv.Append(placeholder)
		if len(g.Foundations[i]) > 0 {
			card := g.Foundations[i][len(g.Foundations[i])-1]
			cardDiv := createDiv(document, attr.Class("card"))
			cardDiv.Value.Set("innerHTML", fmt.Sprintf(`<div class="suit" style="color:%s">%s</div><div class="rank" style="color:%s">%s</div>`, card.Suit.Color(), card.Suit.String(), card.Suit.Color(), card.Rank.String()))
			foundationDiv.Append(cardDiv)
		}
	}

	// Render Tableau
	for i, pile := range g.Tableau {
		pileDiv := ui.Tableau[i]
		pileDiv.Value.Set("innerHTML", "")
		for j, card := range pile {
			var cardAttrs []attr.Attr
			cardAttrs = append(cardAttrs, attr.Class("card"))
			if card == g.SelectedCard {
				cardAttrs = append(cardAttrs, attr.Style("border: 2px solid yellow;"))
			}
			cardAttrs = append(cardAttrs, attr.Style(fmt.Sprintf("top: %dpx;", j*30)))
			if card.FaceUp {
				cardAttrs = append(cardAttrs, attr.Draggable(true))
			}
			cardDiv := createDiv(document, cardAttrs...)
			if card.FaceUp {
				cardDiv.Value.Set("innerHTML", fmt.Sprintf(`<div class="suit" style="color:%s">%s</div><div class="rank" style="color:%s">%s</div>`, card.Suit.Color(), card.Suit.String(), card.Suit.Color(), card.Rank.String()))
			} else {
				cardDiv.SetAttr("style", "background-color: blue;")
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
