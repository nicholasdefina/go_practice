package main

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	charles := person{
		firstName: "Charles",
		lastName:  "Xavier",
		contactInfo: contactInfo{
			email: "chaaaaaaaarles@hotmail.com",
			phone: 18005555555,
		},
	}

	charles.updateFirstName("Jorge")
	charles.print()
}
