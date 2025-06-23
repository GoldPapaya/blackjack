package main

import (
	"fmt"
)

func main() {
	var username string

	fmt.Println("Welcome to the Blackjack Casino! Please enter your name.")
	fmt.Scan(&username)
	// validate username

	fmt.Printf("Welcome %v. Please input a number according to the options below:\n", username)
	startScreen()

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