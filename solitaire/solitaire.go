package main

import (
	"github.com/cshabsin/thegrid/cardkit/ui"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/solitaire/klondike"
)

func main() {
	game := klondike.New()
	document := js.Document()
	boardDiv := document.GetElementByID("game-board")
	ui.NewBoard(game, document, boardDiv)
	select {}
}