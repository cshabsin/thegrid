package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	}
	return ""
}

func (s Suit) Color() string {
	switch s {
	case Clubs, Spades:
		return "black"
	case Diamonds, Hearts:
		return "red"
	}
	return ""
}

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (r Rank) String() string {
	switch r {
	case Ace:
		return "A"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return fmt.Sprintf("%d", r)
	}
}

type Card struct {
	Suit   Suit
	Rank   Rank
	FaceUp bool
}

func NewDeck() []*Card {
	deck := make([]*Card, 0, 52)
	for suit := Clubs; suit <= Spades; suit++ {
		for rank := Ace; rank <= King; rank++ {
			deck = append(deck, &Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

func Shuffle(deck []*Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

type Game struct {
	Deck         []*Card
	Tableau      [7][]*Card
	Foundations  [4][]*Card
	Stock        []*Card
	Waste        []*Card
	SelectedCard *Card
}

func New() *Game {
	deck := NewDeck()
	Shuffle(deck)

	game := &Game{
		Deck: deck,
	}

	// Deal to tableau
	for i := 0; i < 7; i++ { // i is the pile index
		for j := 0; j <= i; j++ { // j is the card index within the pile
			card := game.drawCard()
			if j == i { // Only the last card in the pile is face up
				card.FaceUp = true
			}
			game.Tableau[i] = append(game.Tableau[i], card)
		}
	}

	game.Stock = game.Deck
	game.Deck = nil

	return game
}

func (g *Game) drawCard() *Card {
	card := g.Deck[0]
	g.Deck = g.Deck[1:]
	return card
}

// findAndRemoveSelectedCard finds the selected card, removes it (and any cards on top of it) from its current location,
// and returns the stack of cards that was removed. Returns nil if no card was selected or found.
func (g *Game) findAndRemoveSelectedCard() []*Card {
	if g.SelectedCard == nil {
		return nil
	}

	// Find in tableau
	for i, pile := range g.Tableau {
		for j, card := range pile {
			if card == g.SelectedCard {
				stack := g.Tableau[i][j:]
				g.Tableau[i] = g.Tableau[i][:j]
				return stack
			}
		}
	}

	// Find in waste
	if len(g.Waste) > 0 && g.Waste[len(g.Waste)-1] == g.SelectedCard {
		card := g.Waste[len(g.Waste)-1]
		g.Waste = g.Waste[:len(g.Waste)-1]
		return []*Card{card}
	}

	return nil
}

func (g *Game) MoveSelectedToFoundation(foundationIndex int) {
	if g.SelectedCard == nil {
		return
	}
	foundation := g.Foundations[foundationIndex]
	// Foundation move is only valid if moving a single card.
	// Check if the selected card is the top of its stack.
	if len(g.Waste) > 0 && g.Waste[len(g.Waste)-1] == g.SelectedCard {
		// It's from the waste pile, which is a single card move.
	} else {
		found := false
		for _, pile := range g.Tableau {
			if len(pile) > 0 && pile[len(pile)-1] == g.SelectedCard {
				found = true
				break
			}
		}
		if !found {
			// The selected card is not at the top of a tableau pile, so it's part of a stack. Invalid move.
			return
		}
	}

	if len(foundation) == 0 {
		if g.SelectedCard.Rank != Ace {
			return // Must be an Ace on an empty foundation
		}
	} else {
		topCard := foundation[len(foundation)-1]
		if g.SelectedCard.Suit != topCard.Suit || g.SelectedCard.Rank != topCard.Rank+1 {
			return // Must be same suit and next rank
		}
	}

	stack := g.findAndRemoveSelectedCard()
	if stack != nil {
		g.Foundations[foundationIndex] = append(g.Foundations[foundationIndex], stack...)
		g.SelectedCard = nil
	}
}

func (g *Game) MoveSelectedToTableau(tableauIndex int) {
	if g.SelectedCard == nil {
		return
	}
	destPile := &g.Tableau[tableauIndex]
	if len(*destPile) > 0 {
		topCard := (*destPile)[len(*destPile)-1]
		if g.SelectedCard.Suit.Color() == topCard.Suit.Color() || g.SelectedCard.Rank != topCard.Rank-1 {
			return // Invalid move
		}
	} else {
		if g.SelectedCard.Rank != King {
			return // Must be a king on an empty pile
		}
	}

	stack := g.findAndRemoveSelectedCard()
	if stack != nil {
		*destPile = append(*destPile, stack...)
		g.SelectedCard = nil
	}
}

func (g *Game) CheckWin() bool {
	for _, f := range g.Foundations {
		if len(f) != 13 {
			return false
		}
	}
	return true
}
