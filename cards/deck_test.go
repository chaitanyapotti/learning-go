package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "King Of Diamonds" {
		t.Errorf("Expected the first card to be King of Diamonds, but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()

	d.saveToFile("_decktesting")
	readDeck := newDeckFromFile("_decktesting")

	if len(d) != len(readDeck) {
		t.Errorf("Expected the input %v and output files have the same length of %v", len(d), len(readDeck))
	}
	os.Remove("_decktesting")
}
