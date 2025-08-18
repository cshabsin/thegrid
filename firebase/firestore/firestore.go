package firestore

import (
	"syscall/js"

	"github.com/cshabsin/thegrid/firebase"
)

func InitializeApp(config *firebase.Config) {
	js.Global().Get("theGrid").Get("firebase").Get("firestore").Call("initialize", config.ToJS())
}

// returns a promise?
func CreateGame(game any) js.Value {
	return js.Global().Get("theGrid").Get("firebase").Get("firestore").Call("createGame", js.ValueOf(game))
}
