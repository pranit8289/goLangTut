package main

import "fmt"

func main() {
	fmt.Println("Welcome to Cards...!")
	// var card string = "Ace of Spade" first method to declare variable
	// var card = "Ace of Spade" second method to declare variable
	// card := "Ace of Spade" // third method to declare variable
	// use this at very first initialization of variable
	// var card float = 10.0
	// card := newCard()
	cards := []string{"Ace of spade", newCard()} // declaration of slice one method
	// slice is array but with growing size
	cards = append(cards, "seven of heart") //appending the cars slice
	// append will return new slice, so assigning to cards

	// fmt.Println("cards :", cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Dimonds"
}
