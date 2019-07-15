package main

func main() {
	// card := newCard()
	// cards := newDeck()
	// cards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("cards.txt")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("middle")
	// remainingCards.print()

	newCards := newDeckFromFile("cards.txt")
	newCards.shuffle()
	newCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
