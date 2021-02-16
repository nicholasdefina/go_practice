package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // embeeded struct
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}
