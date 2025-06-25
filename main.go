package main

import (
	"fmt"
	"math/rand"
)

// Card object
type Card struct {
	face string
	suit string
	value int
}

func getHandValue(hand []Card) int {
	total := 0
	numberOfAces := 0
	
	// Calculate total value without aces, and count the number of aces
	for _, card := range hand {
		if card.face == "Ace" {
			numberOfAces += 1
		} else {
			total += card.value
		}
	}

	// Calculate total value of aces and add to total
	for i := 0; i < numberOfAces; i++ {
		if total + 11 <= 21 {
			total += 11
		} else {
			total += 1
		}
	}

	return total
}

// TODO move to its own file
var deck = []Card{
	Card{face: "Ace", suit: "Spades"},
	Card{face: "Ace", suit: "Hearts"},
	Card{face: "Ace", suit: "Diamonds"},
	Card{face: "Ace", suit: "Clubs"},
	Card{face: "Two", suit: "Spades", value: 2},
	Card{face: "Two", suit: "Hearts", value: 2},
	Card{face: "Two", suit: "Diamonds", value: 2},
	Card{face: "Two", suit: "Clubs", value: 2},
	Card{face: "Three", suit: "Spades", value: 3},
	Card{face: "Three", suit: "Hearts", value: 3},
	Card{face: "Three", suit: "Diamonds", value: 3},
	Card{face: "Three", suit: "Clubs", value: 3},
	Card{face: "Four", suit: "Spades", value: 4},
	Card{face: "Four", suit: "Hearts", value: 4},
	Card{face: "Four", suit: "Diamonds", value: 4},
	Card{face: "Four", suit: "Clubs", value: 4},
	Card{face: "Five", suit: "Spades", value: 5},
	Card{face: "Five", suit: "Hearts", value: 5},
	Card{face: "Five", suit: "Diamonds", value: 5},
	Card{face: "Five", suit: "Clubs", value: 5},
	Card{face: "Six", suit: "Spades", value: 6},
	Card{face: "Six", suit: "Hearts", value: 6},
	Card{face: "Six", suit: "Diamonds", value: 6},
	Card{face: "Six", suit: "Clubs", value: 6},
	Card{face: "Seven", suit: "Spades", value: 7},
	Card{face: "Seven", suit: "Hearts", value: 7},
	Card{face: "Seven", suit: "Diamonds", value: 7},
	Card{face: "Seven", suit: "Clubs", value: 7},
	Card{face: "Eight", suit: "Spades", value: 8},
	Card{face: "Eight", suit: "Hearts", value: 8},
	Card{face: "Eight", suit: "Diamonds", value: 8},
	Card{face: "Eight", suit: "Clubs", value: 8},
	Card{face: "Nine", suit: "Spades", value: 9},
	Card{face: "Nine", suit: "Hearts", value: 9},
	Card{face: "Nine", suit: "Diamonds", value: 9},
	Card{face: "Nine", suit: "Clubs", value: 9},
	Card{face: "Ten", suit: "Spades", value: 10},
	Card{face: "Ten", suit: "Hearts", value: 10},
	Card{face: "Ten", suit: "Diamonds", value: 10},
	Card{face: "Ten", suit: "Clubs", value: 10},
	Card{face: "Jack", suit: "Spades", value: 10},
	Card{face: "Jack", suit: "Hearts", value: 10},
	Card{face: "Jack", suit: "Diamonds", value: 10},
	Card{face: "Jack", suit: "Clubs", value: 10},
	Card{face: "Queen", suit: "Spades", value: 10},
	Card{face: "Queen", suit: "Hearts", value: 10},
	Card{face: "Queen", suit: "Diamonds", value: 10},
	Card{face: "Queen", suit: "Clubs", value: 10},
	Card{face: "King", suit: "Spades", value: 10},
	Card{face: "King", suit: "Hearts", value: 10},
	Card{face: "King", suit: "Diamonds", value: 10},
	Card{face: "King", suit: "Clubs", value: 10},
}

func main() {
	var username string

	fmt.Println("Welcome to the Blackjack Casino! Please enter your name.")
	fmt.Scan(&username)
	// validate username

	fmt.Printf("Welcome %v. Please input a number according to the options below:\n", username)
	startScreen()

	// game loop
	for {
		
		// TODO testing stuff below, not final
		var playerCard = Card{face: "King", value: 3}
		var houseCard = Card{face: "Ace"}
		var houseCard2 = Card{face: "Ace"}

		playerHand := []Card{
			playerCard,
		}
		houseHand := []Card{
			houseCard,
			houseCard2,
		}


		fmt.Println("A new round of Blackjack has started.")
		fmt.Printf("You have been dealt an %v, %v", playerCard, getHandValue(playerHand))
		fmt.Printf("The house has an %v, %v", houseCard, getHandValue(houseHand))
		fmt.Println("1 - Hit")
		fmt.Println("2 - Hold")
	}
}

func startScreen() {
	var startScreenInput uint

	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Balance")
	fmt.Println("4 - Ledger")
	fmt.Scan(&startScreenInput)

	switch startScreenInput {
	case 1:
		// play
		fmt.Println("option 1")
	case 2:
		// rules
		fmt.Println("option 2")
		rulesScreen()
	case 3:
		// balance
		fmt.Println("option 3")
		balanceScreen()
	case 4:
		// ledger
		fmt.Println("option 4")
		ledgerScreen()
	}
}

func rulesScreen() {
	// rules
}

func balanceScreen() {
	// balance
}

func ledgerScreen() {
	// ledger
}