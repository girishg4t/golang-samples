//go:generate stringer -type=Suit,Rank

package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
)

type Rank int

const (
	Ace Rank = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	J
	Q
	K
)

var suits = [...]Suit{Spades, Diamonds, Clubs, Hearts}

const (
	minRank = Ace
	mexRank = K
)

//Deck of cards
type Deck struct {
	Suit
	Rank
}

type Cards []Deck

func (d Deck) String() string {
	return d.Suit.String() + " of " + d.Rank.String()
}

func NewDeck() []Deck {
	var cards []Deck
	for _, suit := range suits {
		for rank := minRank; rank <= mexRank; rank++ {
			cards = append(cards, Deck{Suit: suit, Rank: rank})
		}
	}
	return cards
}

//Shuffle the cards
func Shuffle(cards []Deck) []Deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range cards {
		newPosition := r.Intn(len(cards) - 1)
		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
	return cards
}

func (d Cards) Less(i, j int) bool {
	if d[i].Suit == d[j].Suit {
		return d[i].Rank < d[j].Rank
	}
	return d[i].Suit < d[j].Suit
}

func (d Cards) Len() int {
	return len(d)
}
func (d Cards) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

//SortDeck of cards
func SortDeck(cards Cards) {
	sort.Sort(Cards(cards))
}
