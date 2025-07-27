package ui

import (
	"fmt"
	"sort"

	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/pile"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/dragdrop"
	"github.com/cshabsin/thegrid/js/style"
)

type LayoutDirection int

const (
	Horizontal LayoutDirection = iota
	Vertical
)

type PileLayout struct {
	GridRow    int
	GridColumn int
	Direction  LayoutDirection
	CardOffset int
	MaxVisible int
	ClassName  string
}

type Game interface {
	AllCards() []*card.Card
	GetAllPiles() map[string]pile.Pile
	GetPileLayout(name string) PileLayout
	AddListener(func())
	CheckWin() bool
	SelectedCard() *card.Card
	SetSelectedCard(*card.Card)
	MoveToFoundation(*card.Card)
	MoveSelectedToPile(string)
	FlipCard(*card.Card)
	StockClicked()
	ToggleDebugWin()
	NewGame()
}

type Board struct {
	game      Game
	document  js.DOMDocument
	boardDiv  js.DOMElement
	pileToDOM map[string]js.DOMElement
	cardToDOM map[*card.Card]js.DOMElement
}

func (b *Board) reset(doc js.DOMDocument, boardDiv js.DOMElement) {
	boardDiv.Clear()
	b.pileToDOM = make(map[string]js.DOMElement)
	b.cardToDOM = make(map[*card.Card]js.DOMElement)

	// Create all card elements upfront
	for _, c := range b.game.AllCards() {
		b.cardToDOM[c] = b.createCardElement(doc, c)
	}

	// Create pile elements
	var pileNames []string
	for name := range b.game.GetAllPiles() {
		pileNames = append(pileNames, name)
	}
	sort.Strings(pileNames)
	for _, name := range pileNames {
		pileDiv := createDiv(doc, attr.Class("pile"))
		layout := b.game.GetPileLayout(name)
		if layout.ClassName != "" {
			pileDiv.AddClass(layout.ClassName)
		}
		if layout.GridColumn > 0 {
			pileDiv.SetStyle(style.GridColumn(fmt.Sprintf("%d", layout.GridColumn)))
		}
		if layout.GridRow > 0 {
			pileDiv.SetStyle(style.GridRow(fmt.Sprintf("%d", layout.GridRow)))
		}
		b.pileToDOM[name] = pileDiv
		boardDiv.Append(pileDiv)
		func(name string) {
			dragdrop.NewDropTarget(pileDiv, func(e js.DOMEvent) {
				b.game.MoveSelectedToPile(name)
			})
		}(name)
		if name == "stock" {
			pileDiv.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
				b.game.StockClicked()
			})
		}
	}

	newGameButton := createDiv(doc, attr.Class("button"), attr.ID("new-game-button")).SetText("New Game")
	newGameButton.AddEventListener("click", func(_ js.DOMElement, _ js.DOMEvent) {
		b.game.NewGame()
		b.reset(doc, boardDiv)
		b.Render()
	})
	boardDiv.Append(newGameButton)
}

func NewBoard(g Game, doc js.DOMDocument, boardDiv js.DOMElement) *Board {
	b := &Board{
		game:      g,
		document:  doc,
		boardDiv:  boardDiv,
	}
	b.reset(doc, boardDiv)
	g.AddListener(b.Render)
	b.Render()

	doc.AddEventListener("keydown", func(_ js.DOMElement, e js.DOMEvent) {
		if e.Key() == "w" {
			g.ToggleDebugWin()
		}
	})

	return b
}

func (b *Board) Render() {
	// Always remove the win overlay first, we'll add it back if needed.
	winOverlay := b.document.QuerySelector(".win-overlay")
	if !winOverlay.IsNull() {
		winOverlay.Remove()
	}

	for _, c := range b.game.AllCards() {
		cardDiv := b.cardToDOM[c]
		if c == b.game.SelectedCard() {
			cardDiv.AddClass("selected-card")
		} else {
			cardDiv.RemoveClass("selected-card")
		}
	}

	for name, pile := range b.game.GetAllPiles() {
		pileDiv := b.pileToDOM[name]
		pileDiv.Clear()
		layout := b.game.GetPileLayout(name)

		// Dynamically set pile size to encompass all cards for drop targets.
		// Card dimensions are hardcoded for now, matching the CSS.
		cardWidth := 100
		cardHeight := 140
		dropTargetBuffer := 10 // Add a small buffer to the drop target size
		if pile.Len() > 1 {    // Only adjust if there's more than one card
			if layout.Direction == Vertical {
				height := cardHeight + (pile.Len()-1)*layout.CardOffset + dropTargetBuffer
				pileDiv.SetStyle(style.Height(fmt.Sprintf("%dpx", height)))
			} else if layout.Direction == Horizontal {
				width := cardWidth + (pile.Len()-1)*layout.CardOffset + dropTargetBuffer
				pileDiv.SetStyle(style.Width(fmt.Sprintf("%dpx", width)))
			}
		} else {
			// Reset to default size if it was previously larger
			pileDiv.ClearStyles("height", "width")
		}

		if pile.Len() == 0 {
			pileDiv.SetStyle(style.Height("140px"))
			placeholder := createDiv(b.document, attr.Class("card-placeholder"))
			pileDiv.Append(placeholder)
			continue
		}

		start := 0
		if layout.MaxVisible > 0 && pile.Len() > layout.MaxVisible {
			start = pile.Len() - layout.MaxVisible
		}

		for i := start; i < pile.Len(); i++ {
			card := pile[i]
			cardDiv := b.cardToDOM[card]
			resetCardPosition(cardDiv)
			cardDiv.Clear()

			if layout.Direction == Horizontal {
				cardDiv.SetStyle(style.Left(fmt.Sprintf("%dpx", (i-start)*layout.CardOffset)))
			} else {
				cardDiv.SetStyle(style.Top(fmt.Sprintf("%dpx", i*layout.CardOffset)))
			}

			if card.FaceUp {
				cardDiv.SetAttr("draggable", true)
				cardDiv.RemoveClass("face-down-card")
				cardDiv.AddClass("face-up-card")
				populateCardElement(b.document, cardDiv, card)
			} else {
				cardDiv.SetAttr("draggable", false)
				cardDiv.RemoveClass("face-up-card")
				cardDiv.AddClass("face-down-card")
			}
			pileDiv.Append(cardDiv)
		}
	}
	if b.game.CheckWin() {
		overlay := createDiv(b.document, attr.Class("win-overlay"))
		winDiv := createDiv(b.document, attr.Class("win-div"))
		winDiv.Append(b.document.CreateElement("h1").SetText("You Win!"))
		newGameButton := createDiv(b.document, attr.Class("button")).SetText("New Game")
		newGameButton.AddEventListener("click", func(_ js.DOMElement, _ js.DOMEvent) {
			b.game.NewGame()
		})
		winDiv.Append(newGameButton)
		overlay.Append(winDiv)
		b.boardDiv.Append(overlay)
		return
	}
}

func (b *Board) createCardElement(doc js.DOMDocument, c *card.Card) js.DOMElement {
	cardDiv := createDiv(doc, attr.Class("card"))
	dragdrop.NewDraggable(cardDiv, func(e js.DOMEvent) {
		b.game.SetSelectedCard(c)
		cardDiv.AddClass("dragging-source")
	}, func(e js.DOMEvent) {
		cardDiv.RemoveClass("dragging-source")
	})
	cardDiv.AddEventListener("click", func(el js.DOMElement, e js.DOMEvent) {
		if !c.FaceUp {
			b.game.FlipCard(c)
			return
		}
		if b.game.SelectedCard() == c {
			b.game.SetSelectedCard(nil)
		} else {
			b.game.SetSelectedCard(c)
		}
	})
	cardDiv.AddEventListener("dblclick", func(el js.DOMElement, e js.DOMEvent) {
		if !c.FaceUp {
			return
		}
		if g, ok := b.game.(interface {
			MoveToFoundation(*card.Card)
		}); ok {
			g.MoveToFoundation(c)
		}
	})
	return cardDiv
}

func createDiv(doc js.DOMDocument, attrs ...attr.Attr) js.DOMElement {
	return doc.CreateElement("div", attrs...)
}

func populateCardElement(doc js.DOMDocument, cardDiv js.DOMElement, c *card.Card) {
	topSuit := createDiv(doc, attr.Class("suit-top-left")).SetStyle(style.Color(c.Suit.Color()))
	topSuit.SetText(c.Suit.String())
	cardDiv.Append(topSuit)

	topRank := createDiv(doc, attr.Class("rank-top-right")).SetStyle(style.Color(c.Suit.Color()))
	topRank.SetText(c.Rank.String())
	cardDiv.Append(topRank)

	bottomRank := createDiv(doc, attr.Class("rank-bottom-left")).SetStyle(style.Color(c.Suit.Color()))
	bottomRank.SetText(c.Rank.String())
	cardDiv.Append(bottomRank)

	bottomSuit := createDiv(doc, attr.Class("suit-bottom-right")).SetStyle(style.Color(c.Suit.Color()))
	bottomSuit.SetText(c.Suit.String())
	cardDiv.Append(bottomSuit)
}

func resetCardPosition(cardDiv js.DOMElement) {
	cardDiv.ClearStyles("top", "left")
}
