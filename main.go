package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, handTwo := cards.deal(7)
	hand.print()
	handTwo.print()

	fmt.Println(hand.toString())
	fmt.Println(handTwo.toString())

	handWithFunc, handTwoWithFunc := deal(cards, 5)
	handWithFunc.print()
	handTwoWithFunc.print()

	hand.saveToFile("handOfCards")

	newHand := readFromFile("handOfCards")
	fmt.Println("newHand from file", newHand)

	hand.print()
	hand.shuffle()
	hand.print()

}
