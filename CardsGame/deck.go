package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Creating new type 'deck' which is a slice of strings

// All variables of type 'deck' now have access to the methods of type 'deck'
type deck [][2]string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, [2]string{value, suit})
		}
	}

	return cards
}

func (cardsDeck deck) print() {
	for index, card := range cardsDeck {
		fmt.Println(index, card)
	}
}

func deal(cardsDeck deck, handSize int) (deck, deck) {
	hand := cardsDeck[len(cardsDeck)-handSize:]
	cardsDeck = cardsDeck[:len(cardsDeck)-handSize]

	return cardsDeck, hand
}

func toString(cardsDeck deck) string {
	var cardsAsString string
	for index, card := range cardsDeck {
		cardsAsString = cardsAsString + strings.Join(card[:], " of ")

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

	for _, card := range cards {
		cardInfo := strings.Split(card, " of ")
		cardsDeck = append(cardsDeck, [2]string{cardInfo[0], cardInfo[1]})
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
