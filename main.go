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

	nick := person{firstName: "Nick", lastName: "Baughan"}

	nick.firstName = "J Nick"
	nick.contact.email = "email@email.com"
	nick.contact.phone = "(555) 555 5555"
	nick.printName()

	bob := person{
		firstName: "bob",
		lastName:  "dole",
		contact: contact{
			email: "dole@email.com",
			phone: "555 555 5555",
		},
	}
	// & gives you access to the memory address that variable is pointing to
	// pBob := &bob
	// // this will show the memory address
	// println(pBob)
	// not sure why this works... thought i would have to call it off of pBob
	bob.printName() // bob dole
	bob.updateFirstName("bill")
	bob.updateLastName("clinton")
	bob.printName() // bill clinton
}
