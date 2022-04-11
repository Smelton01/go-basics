//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

type Suit uint8
type Rank uint8

type option func([]Card) []Card

// Card represents a playing card
type Card struct {
	Suit Suit
	Rank Rank
}

const (
	minRank = Ace
	maxRank = King
)

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker
)

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

// New produces a new deck of cards with options applied
func New(options ...option) []Card {
	var cards []Card

	for suit := 0; suit <= 3; suit++ {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{
				Suit: Suit(suit),
				Rank: Rank(rank),
			})
		}
	}
	for _, opt := range options {
		cards = opt(cards)
	}
	return cards
}

// WithSortingFunc is an option applied to produce
// a deck sorted according to provided function
func WithSortingFunc(less func(cards []Card) func(i, j int) bool) option {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// WithJokers is an option applied to produce
// a deck including a specified number of joker cards
func WithJokers(count int) option {
	return func(cards []Card) []Card {
		for i := 0; i < count; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

// WithShuffling is an option applied to produce
// a deck of shuffled cards
func WithShuffling(seed int64) option {
	return func(card []Card) []Card {
		rand.Seed(seed)
		rand.Shuffle(len(card), func(i, j int) {
			card[i], card[j] = card[j], card[i]
		})
		return card
	}
}

// WithFilter is an option applied to produce
// a deck excluding cards of specified rank
func WithFilter(ranks ...Rank) option {
	return func(cards []Card) []Card {
		for _, rank := range ranks {
			for idx, card := range cards {
				if card.Rank == rank && card.Suit != Joker {
					cards = append(cards[:idx], cards[idx+1:]...)
				}
			}
		}
		return cards
	}
}

// WithDuplication produces a deck duplicated a specified number of times
func WithDuplication(count int) option {
	return func(cards []Card) []Card {
		finalDeck := cards
		for i := 0; i < count; i++ {
			finalDeck = append(finalDeck, cards...)
		}
		return finalDeck
	}
}

// String uses the command stringer to output formatted representations of cards
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

// DefaultSort provides a default way to sort cards according to rank and suit
func DefaultSort(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
