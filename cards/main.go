package main

import "fmt"

func main() {
	fmt.Println("Welcome to Cards...!")
	// var card string = "Ace of Spade" first method to declare variable
	// var card = "Ace of Spade" second method to declare variable
	// card := "Ace of Spade" // third method to declare variable
	// use this at very first initialization of variable
	// var card float = 10.0
	card := newCard()

	fmt.Println("card :", card)
}

func newCard() string {
	return "Five of Dimonds"
}
