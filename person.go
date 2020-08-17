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
