package main

import (
	"os"
	"testing"
)

//Function to test the nesDeck function
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20,but got %v", len(d))
	}
	if d[0] != "ace of Spades" {
		t.Errorf("Expected first card to be ace of spades bu found %v", d[0])
	}
	if d[len(d)-1] != "five of Clubs" {
		t.Errorf("Expected last card to be five of clubs bu found %v", d[len(d)-1])
	}
}
//function to test the saveTOFile and newDeckFromFile fuctions

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	deck:=newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck:=newDeckFromFile("_decktesting")
	if len(loadedDeck)!=2{
		t.Errorf("Expected deck length of 20,but got %v", len(deck))

	}
	os.Remove("_decktesting")
}