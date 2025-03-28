package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create new type of 'deck'
// which is a slice of strings
type deck []string // deck type extends the definition of slice

// create a new deck of cards
func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{
		"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// this is call 'receiver'
// variable with type 'deck' will have access to this 'print()' method
// 'd' is passed as reference
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal cards
func deal(d deck, handSize int) (deck, deck, error) {
	// throw error if handSize > len(d)
	if handSize > len(d) {
		return nil, nil, fmt.Errorf("handSize of %d is too big for deck of %d", handSize, len(d))
	}
	// return a hand with given handSize and remaining deck
	return d[:handSize], d[handSize:], nil
}

// turn d deck into a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save d deck into a file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// read filename and returns a d deck
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	// err == nil means great!
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	deckString := strings.Split(string(bs), ",")
	d := deck(deckString)
	return d
}

// shuffle a d deck
func (d deck) shuffle() {
	// seed random number generator
	// allow true shuffle everytime the function is called, otherwise it's predictable if rand is used alone.
	// rand.Seed(time.Now().UnixNano()) right now no need to for the purpose of Testings.

	for i := range d {
		// generate rand between (i, len(d)-1)
		j := i + rand.Intn(len(d)-i)
		d[i], d[j] = d[j], d[i]
	}
}
