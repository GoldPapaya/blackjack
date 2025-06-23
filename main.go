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
	startMenu()

}

func startMenu() {
	var startMenuInput uint

	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Balance")
	fmt.Println("4 - Ledger")
	fmt.Scan(&startMenuInput)

	switch startMenuInput {
	case 1:
		// play
		fmt.Println("option 1")
	case 2:
		// rules
		fmt.Println("option 2")
	case 3:
		// balance
		fmt.Println("option 3")
	case 4:
		// ledger
		fmt.Println("option 4")
	}
}