package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var cards = []string{"Five of Diamonds", "Five of Spades", "Five of Hearts", "Five of Clubs"}
	cards = append(cards, "Six of Diamonds")
	cards = append(cards, "Six of Spades")
	cards = append(cards, "Six of Hearts")
	cards = append(cards, "Six of Clubs")

	card := dealNewCard(cards)
	fmt.Println(card)

}

func dealNewCard(deck []string) string {
	min := 0
	max := len(deck) - 1
	rand.Seed(time.Now().UnixNano())
	return deck[rand.Intn(max-min)+min]
}
