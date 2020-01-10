package main

import (
	"fmt"

	d "github.com/girishg4t/golang-samples/quiz/quiz5/deck"
)

func main() {
	card := d.Deck{Suit: 0, Rank: 0}
	fmt.Println(card.String())

	newDeck := d.NewDeck()
	sc := d.Shuffle(newDeck)
	fmt.Println(sc)
	d.SortDeck(sc)
	fmt.Println(sc)
}
