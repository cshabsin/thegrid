package main

import (
	"fmt"

	"github.com/cshabsin/thegrid/cardkit/card"
	"github.com/cshabsin/thegrid/cardkit/ui"
	"github.com/cshabsin/thegrid/js"
)

func main() {
	js.Get("document").Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("Hello, Delta!")
		board := ui.NewBoard(js.Get("document").Call("getElementById", "game-board"), nil)
		board.Draw()
		deck := card.NewDeck(20, 20)
		board.Add(deck)
		return nil
	}))

	select {}
}
