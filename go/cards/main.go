package main

func main() {

	// Shuffle
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// Import & Export
	/*
		cards := newDeck()

		cards.saveToFile("my_cards")

		cards2 := newDeckFromFile("my_cards")
		cards2.print()
	*/
}

func newCard() string {
	return "Five of Diamonds"
}
