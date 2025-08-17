package delta

import "github.com/cshabsin/thegrid/cardkit/card"

type GameStatus string

const (
	GameStatusAcceptingPlayers GameStatus = "still accepting new players"
	GameStatusInProgress      GameStatus = "in progress"
)

type GameState string

const (
	GameStateInitialDeal GameState = "initial deal"
	GameStateAfterFlop   GameState = "after flop"
	GameStateAfterTurn   GameState = "after turn"
	GameStateAfterRiver  GameState = "after river"
)

type Game struct {
	ID            int                  `firestore:"id"`
	Players       []string             `firestore:"players"`
	Status        GameStatus           `firestore:"status"`
	Hands         map[string][]card.Card `firestore:"hands"`
	State         GameState            `firestore:"state"`
	CurrentPlayer string               `firestore:"currentPlayer"`
}
