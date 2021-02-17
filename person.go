package main

import "fmt"

// practice structs

type contactInfo struct {
	email string
	phone int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // embeeded struct
}

func createPerson(first string, last string, email string, number int) person {
	return person{
		firstName: first,
		lastName:  last,
		contactInfo: contactInfo{
			email: email,
			phone: number,
		},
	}
}

// * in front of type completely different than * in front of variable
func (pointerToPerson *person) updateFirstName(newFirstName string) { // * person: pointer that points to type person (or a person)
	(*pointerToPerson).firstName = newFirstName // * operator: give me value this memory address is pointing at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
