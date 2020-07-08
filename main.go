package main

import "fmt"

func main() {
	cards := newDeck()
	hand, handTwo := cards.deal(3)
	hand.print()
	handTwo.print()

	fmt.Println(hand.toString())
	fmt.Println(handTwo.toString())

	handWithFunc, handTwoWithFunc := deal(cards, 5)
	handWithFunc.print()
	handTwoWithFunc.print()

	hand.saveToFile("handOfCards")

}
