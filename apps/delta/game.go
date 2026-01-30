package delta

import "github.com/cshabsin/thegrid/cardkit/card"

type GameStatus int

const (
	GameStatusUnknown GameStatus = iota
	GameStatusAcceptingPlayers
	GameStatusInProgress
)

func (s GameStatus) String() string {
	switch s {
	case GameStatusAcceptingPlayers:
		return "still accepting new players"
	case GameStatusInProgress:
		return "in progress"
	default:
		return "unknown"
	}
}

type GameState int

const (
	GameStateUnknown GameState = iota
	GameStateInitialDeal
	GameStateAfterFlop
	GameStateAfterTurn
	GameStateAfterRiver
)

func (s GameState) String() string {
	switch s {
	case GameStateInitialDeal:
		return "initial deal"
	case GameStateAfterFlop:
		return "after flop"
	case GameStateAfterTurn:
		return "after turn"
	case GameStateAfterRiver:
		return "after river"
	default:
		return "unknown"
	}
}

type Game struct {
	ID            int                  `firestore:"id"`
	Players       []string             `firestore:"players"`
	Status        GameStatus           `firestore:"status"`
	Hands         map[string][]card.Card `firestore:"hands"`
	State         GameState            `firestore:"state"`
	CurrentPlayer string               `firestore:"currentPlayer"`
}