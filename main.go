package main

import (
	"fmt"
)

// Card object
type Card struct {
	face string
	value uint
}

func getHandValue(hand []Card) uint {
	var total uint = 0
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