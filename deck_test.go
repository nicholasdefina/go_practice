package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got: %v", len(d))
	}

	firstCard := d[0]
	if firstCard.suit != "Spades" || firstCard.value != "Ace" {
		t.Errorf("Expected first card to be Ace of Spades but got: %v", firstCard)
	}
	lastCard := d[len(d)-1]
	if lastCard.suit != "Diamonds" || lastCard.value != "King" {
		t.Errorf("Expected first card to be King of Diamonds but got: %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_deckTesting"
	// cleanup in case of existing file
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length of 52 but got: %v", len(loadedDeck))
	}

	os.Remove(fileName)

}
