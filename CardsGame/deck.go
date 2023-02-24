package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Creating new type 'deck' which is a slice of strings

// All variables of type 'deck' now have access to the methods of type 'deck'
type card struct {
	value string
	suit  string
}
type deck []card

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, card{value, suit})
		}
	}

	return cards
}

func (cardsDeck deck) print() {
	for index, _card := range cardsDeck {
		fmt.Println(index, _card)
	}
}

func (deckPointer *deck) deal(handSize int) deck {
	hand := (*deckPointer)[len((*deckPointer))-handSize:]
	*deckPointer = (*deckPointer)[:len((*deckPointer))-handSize]

	return hand
}

func toString(cardsDeck deck) string {
	var cardsAsString string
	for index, _card := range cardsDeck {
		cardsAsString = cardsAsString + _card.value + " of " + _card.suit

		if index < len(cardsDeck)-1 {
			cardsAsString = cardsAsString + ","
		}
	}
	return cardsAsString
}

func (cardsDeck deck) saveToFile(filename string) error {
	lastIndex := strings.LastIndex(filename, "/")
	err := os.MkdirAll(filename[:lastIndex], os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(toString(cardsDeck)), os.ModePerm)
}

func newDeckFromFile(filename string) deck {
	var cardsDeck deck

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cards := strings.Split(string(data[:]), ",")

	for _, _card := range cards {
		cardInfo := strings.Split(_card, " of ")
		cardsDeck = append(cardsDeck, card{cardInfo[0], cardInfo[1]})
	}

	return cardsDeck
}

func (cardsDeck deck) shuffle() {
	rand.Shuffle(len(cardsDeck), func(i, j int) {
		cardsDeck[i], cardsDeck[j] = cardsDeck[j], cardsDeck[i]
	})
}

func (cardsDeck deck) isEqual(secondCardsDeck deck) bool {
	if len(cardsDeck) != len(secondCardsDeck) {
		return false
	}
	for index, card := range cardsDeck {
		if card != secondCardsDeck[index] {
			return false
		}
	}
	return true
}
