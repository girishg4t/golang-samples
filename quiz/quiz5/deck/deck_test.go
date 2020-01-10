package deck

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

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
