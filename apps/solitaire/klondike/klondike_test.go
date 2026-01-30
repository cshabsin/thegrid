package klondike

import (
	"testing"

	"github.com/cshabsin/thegrid/cardkit/card"
)

func TestNewGame(t *testing.T) {
	game := New()
	if len(game.Stock) != 52-28 {
		t.Errorf("Expected 24 cards in stock, got %d", len(game.Stock))
	}
	for i, pile := range game.Tableau {
		if len(pile) != i+1 {
			t.Errorf("Expected %d cards in tableau pile %d, got %d", i+1, i, len(pile))
		}
	}
}

func TestMoveToTableau(t *testing.T) {
	game := &Klondike{}
	game.Tableau[0] = []*card.Card{
		{Suit: card.Spades, Rank: card.King, FaceUp: true},
	}
	game.Tableau[1] = []*card.Card{
		{Suit: card.Hearts, Rank: card.Queen, FaceUp: true},
	}
	game.SelectedCard = game.Tableau[1][0]
	game.MoveSelectedToTableau(0)
	if len(game.Tableau[0]) != 2 {
		t.Errorf("Expected 2 cards in tableau pile 0, got %d", len(game.Tableau[0]))
	}
	if len(game.Tableau[1]) != 0 {
		t.Errorf("Expected 0 cards in tableau pile 1, got %d", len(game.Tableau[1]))
	}
	if game.SelectedCard != nil {
		t.Error("Expected selected card to be nil")
	}
}

func TestInvalidMoveToTableau(t *testing.T) {
	game := &Klondike{}
	game.Tableau[0] = []*card.Card{
		{Suit: card.Spades, Rank: card.King, FaceUp: true},
	}
	game.Tableau[1] = []*card.Card{
		{Suit: card.Spades, Rank: card.Queen, FaceUp: true},
	}
	game.SelectedCard = game.Tableau[1][0]
	game.MoveSelectedToTableau(0)
	if len(game.Tableau[0]) != 1 {
		t.Errorf("Expected 1 card in tableau pile 0, got %d", len(game.Tableau[0]))
	}
	if len(game.Tableau[1]) != 1 {
		t.Errorf("Expected 1 card in tableau pile 1, got %d", len(game.Tableau[1]))
	}
	if game.SelectedCard == nil {
		t.Error("Expected selected card to not be nil")
	}
}

func TestMoveToFoundation(t *testing.T) {
	game := &Klondike{}
	game.Tableau[0] = []*card.Card{
		{Suit: card.Spades, Rank: card.Ace, FaceUp: true},
	}
	game.SelectedCard = game.Tableau[0][0]
	game.MoveSelectedToFoundation(0)
	if len(game.Tableau[0]) != 0 {
		t.Errorf("Expected 0 cards in tableau pile 0, got %d", len(game.Tableau[0]))
	}
	if len(game.Foundations[0]) != 1 {
		t.Errorf("Expected 1 card in foundation 0, got %d", len(game.Foundations[0]))
	}
	if game.SelectedCard != nil {
		t.Error("Expected selected card to be nil")
	}
}

func TestInvalidMoveToFoundation(t *testing.T) {
	game := &Klondike{}
	game.Foundations[0] = []*card.Card{
		{Suit: card.Spades, Rank: card.Ace, FaceUp: true},
	}
	game.Tableau[0] = []*card.Card{
		{Suit: card.Spades, Rank: card.Three, FaceUp: true},
	}
	game.SelectedCard = game.Tableau[0][0]
	game.MoveSelectedToFoundation(0)
	if len(game.Tableau[0]) != 1 {
		t.Errorf("Expected 1 card in tableau pile 0, got %d", len(game.Tableau[0]))
	}
	if len(game.Foundations[0]) != 1 {
		t.Errorf("Expected 1 card in foundation 0, got %d", len(game.Foundations[0]))
	}
	if game.SelectedCard == nil {
		t.Error("Expected selected card to not be nil")
	}
}
