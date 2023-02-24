package main

import (
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestNewDeck(handler *testing.T) {
	cardsDeck := newDeck()

	if len(cardsDeck) != 52 {
		handler.Errorf("Expected 'cardsDeck' to have 52 cards but counted %v", len(cardsDeck))
	}

	if cardsDeck[0] != [2]string{"A", "Spades"} {
		handler.Errorf("Expected the first card 'cardsDeck' to be [A Spades] but got %v", cardsDeck[0])
	}

	if cardsDeck[len(cardsDeck)-1] != [2]string{"K", "Clubs"} {
		handler.Errorf(
			"Expected the first card 'cardsDeck' to be [K Clubs] but got %v",
			cardsDeck[len(cardsDeck)-1],
		)
	}
}

func TestSaveToFileAndNewDeckFromFile(handler *testing.T) {
	filename := "Test Deck/test_deck.txt"
	cardsDeck := newDeck()
	cardsDeck.shuffle()

	err := cardsDeck.saveToFile(filename)
	if err != nil {
		handler.Errorf("A unexpected error happened.  Error: %v", err)
	}
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		handler.Fatalf("File '%v' was not found.", filename)
	}

	cardsDeckFromFile := newDeckFromFile(filename)
	if !cardsDeckFromFile.isEqual(cardsDeck) {
		handler.Errorf(
			"Cards in deck retrieved from file '%v' are not in the same order as the original deck created.",
			filename,
		)
	}

	// Clean up files
	err = os.RemoveAll(filename)
	if err != nil {
		handler.Errorf("File '%v' could not be removed.", filename)
	}
}

func TestShuffle(handler *testing.T) {
	cardsDeck := newDeck()
	constantCardsDeck := newDeck()
	cardsDeck.shuffle()

	if constantCardsDeck.isEqual(cardsDeck) {
		handler.Fatalf("Deck of cards should be of the same length but different ordering.")
	}
}

func TestToString(handler *testing.T) {
	cardsDeck := newDeck()
	cardsDeckString := toString(cardsDeck)

	if reflect.TypeOf(cardsDeckString).Name() != "string" {
		handler.Fatalf(
			"Expected deck of cards in string format but got '%v' format instead.",
			reflect.TypeOf(cardsDeckString).Name(),
		)
	}
}

func TestIsEqual(handler *testing.T) {
	firstCardsDeck := newDeck()
	secondCardsDeck := newDeck()

	if reflect.TypeOf(firstCardsDeck.isEqual(secondCardsDeck)).Name() != "bool" {
		handler.Fatalf(
			"Expected return in 'bool' format but got '%v' format instead.",
			reflect.TypeOf(firstCardsDeck.isEqual(secondCardsDeck)).Name(),
		)
	}
}

func TestDeal(handler *testing.T) {
	cardsDeck := newDeck()
	originalDeckLength := len(cardsDeck)
	handSize := rand.Intn(originalDeckLength)

	cardsDeck, hand := deal(cardsDeck, handSize)

	if len(cardsDeck) != originalDeckLength-handSize {
		handler.Errorf(
			"Expected 'cardsDeck' length of '%v' but got '%v'.",
			originalDeckLength-handSize,
			len(cardsDeck),
		)
	}
	if len(hand) != handSize {
		handler.Errorf(
			"Expected 'hand' length of '%v' but got '%v'.",
			handSize,
			len(hand),
		)
	}
}
