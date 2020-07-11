package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Queen", "King", "Jack"}
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

func (d deck) toString() string {
	// since a deck "extends" []strings we can easily convert
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	cards := deck{}
	cSlice := strings.Split(string(bSlice), ",")

	for _, card := range cSlice {
		cards = append(cards, card)
	}
	return cards

}

// is there an advantage to using a receiver function vs a function???
// both seem to work
func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize : handSize+handSize]
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize : handSize+handSize]
}
