package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create type 'deck,' a slice of cards
type card struct {
	value string
	suit  string
}

type deck []card

func newDeck() deck {
	d := deck{}
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			d = append(d, card{value: value,
				suit: suit,
			})
		}
	}
	return d
}

// 'receiver' comes before func name.
func (d deck) print() { // (d deck) is the 'receiver' to the function. 'd' is similar to 'self' in python.
	for idx, card := range d { //  convention is one letter abbrev of type. any variable of type deck gets access to print method
		fmt.Println(idx, card)
	}

}

func deal(d deck, numOfCards int) (deck, deck) {
	return d[:numOfCards], d[numOfCards:]
}

func (d deck) toString() string {
	var deckString []string
	for _, val := range d {
		deckString = append(deckString, string(val.value+" of "+val.suit))
	}
	return strings.Join([]string(deckString), ",") // make slice of type string, join to comma-sep string
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 anyone can read/write this file
}

func newDeckFromFile(filename string) deck {
	byteStr, err := ioutil.ReadFile(filename)

	if err != nil {
		// error!
		fmt.Println("Error encountered in newDeckFromFile: ", err)
		os.Exit(1) // any non-zero is not a success
	}

	deckString := strings.Split(string(byteStr), ",")
	newDeck := deck{}
	for _, val := range deckString {
		c := strings.Split(val, " of ")
		newDeck = append(newDeck, card{value: c[0], suit: c[1]})

	}
	return newDeck
}

func (d deck) shuffle() {
	// Go rand.Intn uses same seed for random generation, need to adjust seed.
	seed := time.Now().UnixNano()  // random seed via tim. unix nano returns int64
	source := rand.NewSource(seed) // create new seed source
	r := rand.New(source)          // create new rand type

	for idx := range d {
		newPosition := r.Intn(len(d) - 1)
		d[idx], d[newPosition] = d[newPosition], d[idx] // one liner swap, like python3
	}
}
