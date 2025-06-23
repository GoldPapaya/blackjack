package main

import (
	"fmt"
)

func main() {
	var username string

	fmt.Println("Welcome to the Blackjack Casino! Please enter your name.")
	fmt.Scan(&username)
	// validate username

	var startScreenInput uint

	fmt.Printf("Welcome %v. Please input a number according to the options below:\n", username)
	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Balance")
	fmt.Println("4 - Ledger")
	fmt.Scan(&startScreenInput)

	fmt.Printf("%v", startScreenInput)
}