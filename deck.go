package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Queen", "King", "Jack"}
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// is there an advantage to using a receiver function vs a function???
// both seem to work
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize : handSize+handSize]
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize : handSize+handSize]
}
