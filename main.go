package main

func main() {
	cards := getDeckFromFile("cards")
	// cards.print()
	cards.shuffle()
	cards.print()
}
