package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v", len(d))
	}

	if d[0] != "Ace of S" {
		t.Errorf("Expect but got %v", d[0])
	}

	if d[len(d)-1] != "Four of C" {
		t.Errorf("Expecte but got %v ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
