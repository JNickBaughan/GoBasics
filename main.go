package main

func main() {
	cards := deck{"Five of Diamonds", "Five of Spades", "Five of Hearts", "Five of Clubs"}
	cards.print()
	cards.deal()
}
