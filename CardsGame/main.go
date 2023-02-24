package main

func main() {
	// 2 Different ways of defining a variable.  Type of variable will be inferred from right-hand side if it is not
	// declared.

	// Way #1:
	//var card string = "Ace of Spades"

	// Way #2:
	// The `:=` means that it will create and assign a value to a variable.
	//card := "Ace of Spades"

	// Way #3:
	// Initialize variable first, then later assign value to it.
	//var card string
	//card = "Ace of Spades"

	// Way #4:
	// Define the variable outside `func main()`.
	// Assign a value to the variable inside `func main()`.
	//var card string
	//card = "Ace of Spades"

	deckOfCards := newDeck()
	deckOfCards.shuffle()
	deckOfCards.print()
}
