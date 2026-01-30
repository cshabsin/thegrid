package main

import (
	"github.com/cshabsin/thegrid/apps/solitaire/klondike"
	"github.com/cshabsin/thegrid/cardkit/ui"
	"github.com/cshabsin/thegrid/js"
)

func main() {
	game := klondike.New()
	document := js.Document()
	boardDiv := document.GetElementByID("game-board")
	ui.NewBoard(game, document, boardDiv)
	select {}
}
