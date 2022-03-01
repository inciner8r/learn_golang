package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("deck length is %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("first card is %v", d[0])
	}
	if d[len(d)-1] != "Four of Clover" {
		t.Errorf("last card is %v", d[len(d)-1])
	}
}
