package deck

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Hearts})
	fmt.Println(Card{Rank: Two, Suit: Spades})
	fmt.Println(Card{Rank: Nine, Suit: Diamonds})
	fmt.Println(Card{Rank: Jack, Suit: Clubs})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(WithSortingFunc(DefaultSort))
	want := Card{Rank: Ace, Suit: Spades}
	if cards[0] != want {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestWithJokers(t *testing.T) {
	cards := New(WithJokers(4))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 4 {
		t.Error("Expected 4 Jokers, received:", count)
	}
}

func TestWithShuffling(t *testing.T) {
	seed := time.Now().Unix()

	got := New(WithShuffling(seed))

	want := New()
	rand.Seed(seed)
	rand.Shuffle(len(want), func(i, j int) { want[i], want[j] = want[j], want[i] })

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("Expected %v, got %v", want[i], got[i])
		}
	}
}

func TestWithDuplication(t *testing.T) {
	cards := New(WithDuplication(3))

	want := 13 * 4 * 4
	if len(cards) != want {
		t.Errorf("Expected %v card deck, got %v cards.", want, len(cards))
	}
}

func TestWithFilter(t *testing.T) {
	filterRank := Two

	cards := New(WithFilter(filterRank))

	for _, card := range cards {
		if card.Rank == filterRank {
			t.Errorf("Expected Rank %s to be filtered out, %s found in deck.", card.Rank, card.String())
		}
	}
}
