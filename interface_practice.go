package main

import (
	"fmt"
)

// interfaces function a little like inheritance/polymorphism for python/oop. interfaces are an abstract type, cannot create
// direct instance of a bot. interfaces are implicit, no need to explicitly state children that are 'bots'
type bot interface {
	getGreeting() string // any types that have all funcs and matching return types is member of this bot type
	// if another func was here, children would have to implement this func as well in order to be a 'bot'
}

// this function now available to types that are members of the bot type
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

func (e englishBot) getGreeting() string {
	return "Hello!"
}

func (s spanishBot) getGreeting() string {
	return "Hola!"
}

// custom writer
type logger struct{}

func (logger) Write(byteSlice []byte) (int, error) {
	fmt.Println(string(byteSlice))
	fmt.Println("just wrote", len(byteSlice), "bytes!")
	return len(byteSlice), nil
}


// shapes
type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type triangle struct{
	base float64
	height float64
}

type square struct{
	length float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}
func (s square) getArea() float64 {
	return s.length * s.length
}
