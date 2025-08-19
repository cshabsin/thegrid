package firestore

import (
	"syscall/js"

	"github.com/cshabsin/thegrid/js"
)

func InitializeApp(app js.Value) {
	js.Global().Get("theGrid").Get("firebase").Get("firestore").Call("initialize", app)
}

func CreateGame(game any) js.Promise {
	return js.Promise{Value: js.Global().Get("theGrid").Get("firebase").Get("firestore").Call("createGame", js.ValueOf(game))}
}