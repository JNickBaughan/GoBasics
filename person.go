package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contact   contact
}

// * is an operator that gets you the value at the memory address
func (pointerToPerson *person) printName() {
	fmt.Println((*pointerToPerson).firstName + " " + (*pointerToPerson).lastName)
}

// when there is a *in front of a type it means it is a pointer to that type, not the same as an operator
// inside the receiver the * pointerToPerson turns it back from a pointer to a value
// this will actually accept a type person OR a pointer to a person
func (pointerToPerson *person) printEmail() {
	fmt.Println((*pointerToPerson).firstName + " " + (*pointerToPerson).lastName + " => email is " + (*pointerToPerson).contact.email)
}

func (pointerToPerson *person) printPhone() {
	fmt.Println((*pointerToPerson).firstName + " " + (*pointerToPerson).lastName + " => phone is " + (*pointerToPerson).contact.phone)
}

func (pointerToPerson *person) updateFirstName(firstNameUpdate string) {
	(*pointerToPerson).firstName = firstNameUpdate
}

func (pointerToPerson *person) updateLastName(lastNameUpdate string) {
	(*pointerToPerson).lastName = lastNameUpdate
}
