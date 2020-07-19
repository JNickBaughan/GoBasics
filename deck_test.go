package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	length := len(deck)
	first := deck[0]
	last := deck[length-1]
	if length != 52 {
		t.Errorf("Expected deck length of 16, but got %v", length)
	}

	if first != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, but got %v", first)
	}

	if last != "Jack of Clubs" {
		t.Errorf("Expected the first card to be Jack of Club, but got %v", last)
	}
}
