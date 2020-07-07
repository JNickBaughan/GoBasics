package main

func main() {
	cards := newDeck()
	hand, handTwo := cards.deal(3)
	hand.print()
	handTwo.print()

	handWithFunc, handTwoWithFunc := deal(cards, 5)
	handWithFunc.print()
	handTwoWithFunc.print()

}
