package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := dealCards(cards, 5)

	hand.PrintCards()
	remainingCards.PrintCards()

	fmt.Println(cards.toString())

	cards.shuffleCards()

	cards.PrintCards()

	cards.saveToFile("myCards")

	cards = newDeckFromFile("myCards")

	cards.PrintCards()
}
