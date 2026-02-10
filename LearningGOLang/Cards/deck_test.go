package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expeced first card is Ace of spades, but actually got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expeced first card is King of Clubs, but actually got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_deckTesting"
	os.Remove(filename)
	deck := newDeck()
	deck.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(deck))
	}
	os.Remove(filename)
}
