package main

import (
	"fmt"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/solitaire/game"
)

type GameUI struct {
	Board        js.DOMElement
	Stock        js.DOMElement
	Waste        js.DOMElement
	Foundations  [4]js.DOMElement
	Tableau      [7]js.DOMElement
}

var g *game.Game

func createDivWithStyles(doc js.DOMDocument, styles map[string]string) js.DOMElement {
	div := doc.CreateElement("div")
	for key, value := range styles {
		div.Value.Get("style").Set(key, value)
	}
	return div
}

func main() {
	g = game.New()
	document := js.Document()
	board := document.GetElementByID("game-board")
	board.Value.Set("innerHTML", "") // Clear the board

	ui := &GameUI{Board: board}

	// Create top row elements
	topRow := createDivWithStyles(document, map[string]string{
		"gridRow":             "1",
		"gridColumn":          "1 / span 7",
		"display":             "grid",
		"gridTemplateColumns": "repeat(7, 1fr)",
	})
	board.Append(topRow)

	ui.Stock = createDivWithStyles(document, map[string]string{
		"gridColumn": "1",
		"position":   "relative",
	})
	topRow.Append(ui.Stock)

	ui.Waste = createDivWithStyles(document, map[string]string{
		"gridColumn": "2",
		"position":   "relative",
	})
	topRow.Append(ui.Waste)
	ui.Waste.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if len(g.Waste) > 0 {
			g.SelectedCard = g.Waste[len(g.Waste)-1]
			render(document, ui, g)
		}
	})

	for i := 0; i < 4; i++ {
		ui.Foundations[i] = createDivWithStyles(document, map[string]string{
			"gridColumn": fmt.Sprintf("%d", i+4),
			"position":   "relative",
		})
		topRow.Append(ui.Foundations[i])
		foundationIndex := i
		ui.Foundations[i].AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
			if g.SelectedCard == nil {
				return
			}
			foundation := g.Foundations[foundationIndex]
			if len(foundation) == 0 {
				if g.SelectedCard.Rank == game.Ace {
					g.MoveSelectedToFoundation(foundationIndex)
					render(document, ui, g)
				}
				return
			}
			topCard := foundation[len(foundation)-1]
			if g.SelectedCard.Suit == topCard.Suit && g.SelectedCard.Rank == topCard.Rank+1 {
				g.MoveSelectedToFoundation(foundationIndex)
				render(document, ui, g)
			}
		})
	}

	// Create tableau elements
	tableauRow := createDivWithStyles(document, map[string]string{
		"gridRow":             "2",
		"gridColumn":          "1 / span 7",
		"display":             "grid",
		"gridTemplateColumns": "repeat(7, 1fr)",
	})
	board.Append(tableauRow)
	for i := 0; i < 7; i++ {
		ui.Tableau[i] = createDivWithStyles(document, map[string]string{
			"gridColumn": fmt.Sprintf("%d", i+1),
			"position":   "relative",
		})
		tableauRow.Append(ui.Tableau[i])
		pileIndex := i
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
								render(document, ui, g)
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
			render(document, ui, g)
		})
	}

	ui.Stock.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if len(g.Stock) > 0 {
			// Draw up to 3 cards
			numToDraw := 3
			if len(g.Stock) < 3 {
				numToDraw = len(g.Stock)
			}
			for i := 0; i < numToDraw; i++ {
				card := g.Stock[0]
				g.Stock = g.Stock[1:]
				card.FaceUp = true
				g.Waste = append(g.Waste, card)
			}
		} else if len(g.Waste) > 0 {
			g.Stock = g.Waste
			g.Waste = nil
			for _, card := range g.Stock {
				card.FaceUp = false
			}
		}
		render(document, ui, g)
	})

	render(document, ui, g)
	select {}
}

func render(document js.DOMDocument, ui *GameUI, g *game.Game) {
	if g.CheckWin() {
		winDiv := createDivWithStyles(document, map[string]string{
			"position":  "absolute",
			"top":       "50%",
			"left":      "50%",
			"transform": "translate(-50%, -50%)",
			"fontSize":  "5em",
			"color":     "white",
		})
		winDiv.Value.Set("innerHTML", "<h1>You Win!</h1>")
		ui.Board.Append(winDiv)
		return
	}
	// Render Stock
	ui.Stock.Value.Set("innerHTML", "")
	stockCardDiv := document.CreateElement("div")
	stockCardDiv.Value.Set("className", "card")
	if len(g.Stock) > 0 {
		stockCardDiv.Value.Get("style").Set("backgroundColor", "blue")
	} else {
		stockCardDiv.Value.Get("style").Set("border", "1px solid black")
	}
	ui.Stock.Append(stockCardDiv)

	// Render Waste
	ui.Waste.Value.Set("innerHTML", "")
	wastePlaceholder := document.CreateElement("div")
	wastePlaceholder.Value.Set("className", "card")
	wastePlaceholder.Value.Get("style").Set("border", "1px solid black")
	ui.Waste.Append(wastePlaceholder)
	if len(g.Waste) > 0 {
		start := len(g.Waste) - 3
		if start < 0 {
			start = 0
		}
		for i := start; i < len(g.Waste); i++ {
			card := g.Waste[i]
			cardDiv := document.CreateElement("div")
			cardDiv.Value.Set("className", "card")
			// The top card of the waste is the one that can be selected.
			if i == len(g.Waste)-1 && card == g.SelectedCard {
				cardDiv.Value.Get("style").Set("border", "2px solid yellow")
			}
			// Offset the cards to fan them out
			cardDiv.Value.Get("style").Set("left", fmt.Sprintf("%dpx", (i-start)*20))
			cardDiv.Value.Set("innerHTML", fmt.Sprintf(`<div class="suit" style="color:%s">%s</div><div class="rank" style="color:%s">%s</div>`, card.Suit.Color(), card.Suit.String(), card.Suit.Color(), card.Rank.String()))
			ui.Waste.Append(cardDiv)
			cardDiv.AddEventListener("dblclick", func(el js.DOMElement, e js.DOMEvent) {
				for i := 0; i < 4; i++ {
					foundation := g.Foundations[i]
					if len(foundation) == 0 {
						if card.Rank == game.Ace {
							g.SelectedCard = card
							g.MoveSelectedToFoundation(i)
							render(document, ui, g)
							return
						}
					} else {
						topCard := foundation[len(foundation)-1]
						if card.Suit == topCard.Suit && card.Rank == topCard.Rank+1 {
							g.SelectedCard = card
							g.MoveSelectedToFoundation(i)
							render(document, ui, g)
							return
						}
					}
				}
			})
		}
	}

	// Render Foundations
	for i := 0; i < 4; i++ {
		foundationDiv := ui.Foundations[i]
		foundationDiv.Value.Set("innerHTML", "")
		placeholder := document.CreateElement("div")
		placeholder.Value.Set("className", "card")
		placeholder.Value.Get("style").Set("border", "1px solid black")
		foundationDiv.Append(placeholder)
		if len(g.Foundations[i]) > 0 {
			card := g.Foundations[i][len(g.Foundations[i])-1]
			cardDiv := document.CreateElement("div")
			cardDiv.Value.Set("className", "card")
			cardDiv.Value.Set("innerHTML", fmt.Sprintf(`<div class="suit" style="color:%s">%s</div><div class="rank" style="color:%s">%s</div>`, card.Suit.Color(), card.Suit.String(), card.Suit.Color(), card.Rank.String()))
			foundationDiv.Append(cardDiv)
		}
	}

	// Render Tableau
	for i, pile := range g.Tableau {
		pileDiv := ui.Tableau[i]
		pileDiv.Value.Set("innerHTML", "")
		for j, card := range pile {
			cardDiv := document.CreateElement("div")
			cardDiv.Value.Set("className", "card")
			if card == g.SelectedCard {
				cardDiv.Value.Get("style").Set("border", "2px solid yellow")
			}
			cardDiv.Value.Get("style").Set("top", fmt.Sprintf("%dpx", j*30))
			if card.FaceUp {
				cardDiv.Value.Set("innerHTML", fmt.Sprintf(`<div class="suit" style="color:%s">%s</div><div class="rank" style="color:%s">%s</div>`, card.Suit.Color(), card.Suit.String(), card.Suit.Color(), card.Rank.String()))
			} else {
				cardDiv.Value.Get("style").Set("backgroundColor", "blue")
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
							render(document, ui, g)
							return
						}
					} else {
						topCard := foundation[len(foundation)-1]
						if card.Suit == topCard.Suit && card.Rank == topCard.Rank+1 {
							g.SelectedCard = card
							g.MoveSelectedToFoundation(i)
							render(document, ui, g)
							return
						}
					}
				}
			})
		}
	}
}
