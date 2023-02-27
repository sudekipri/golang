package main

// import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	cards.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()
}