package main

import "log"

func main() {
	// var card string = "Ace of Spades"
	// := is needed to initialize a variable, no need for re-assign
	// card := "Aces of Spades"
	// card = "Five of Diamonds"

	// card := newCard()
	// fmt.Println(card)

	// this is a slice (fixable) not array. Works just like Java, Cpp
	// cards := []string{"Ace of Diamonds", newCard()}
	cards := newDeck()
	// cards.print()

	hand, _, err := deal(cards, 13)
	if err != nil {
		log.Println("Split failed:", err)
	}
	hand.print()
	hand.shuffle()
	hand.print()
	hand.saveToFile("my_hand")
	// remainingHand.print()
	// cards.saveToFile("my_cards")
	// cards = newDeckFromFile("my_cards")
	// cards.print()
}
