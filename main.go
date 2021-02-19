package main

func main() {
	// // classic deck o' cards
	// cards := newDeck()
	// cards.shuffle()

	// // simple structs
	// charles := createPerson("Charles", "Xavier", "charrruls@hotmail.com", 18005555555)
	// charles.updateFirstName("Jorge") // can shorthand the pointer like this. receiver function can accept pointer to person or a person
	// charles.print()

	// // simple maps
	// colors := createColorsMap()
	// fmt.Println(colors)
	// printColors(colors)

	// // simple interfaces
	// eb := englishBot{}
	// sb := spanishBot{}
	// printGreeting(eb)
	// printGreeting(sb)

	// t := triangle{base: 3,
	// 	height: 6}
	// s := square{length: 5}
	// printArea(t)
	// printArea(s)

	// // http requests in go
	// resp, err := http.Get("https://google.com")
	// if err != nil {
	// 	fmt.Println("Error was:", err)
	// 	os.Exit(1)
	// }
	// io.Copy(logger{}, resp.Body) // io.Copy needs 1st arg that uses Writer interface (custom write here, could also use os.Stdout), 2nd that uses reader

	siteChecker()
}
