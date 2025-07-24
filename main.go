package main

import (
	"fmt"
	"os"
)

var deck []Card
var balance int = 1000
var ledger []int

func playerTurn(deck []Card, playerHand []Card, houseHand []Card, bet int) int {
	var gameInput int
	for {
		clearCLI()
		if getHandValue(playerHand) > 21 {
			fmt.Printf("You drew a %v, increasing your hand value to %v.\n", playerHand[len(playerHand)-1], getHandValue(playerHand))
			fmt.Println("Your hand went bust!")
			fmt.Println("Enter 1 to continue...")
			fmt.Scan(&gameInput)
			adjustBalance(2, bet)
			break
		} else if getHandValue(playerHand) == 21 {
			clearCLI()
			fmt.Printf("You were dealt a %v.\n", playerHand[len(playerHand)-1])
			fmt.Println("Since your hand has a value of 21, you stand.")
			fmt.Printf("The house reveals their face down card to be a %v.\n", houseHand[1])
			houseTurn(deck, playerHand, houseHand, bet)
			return 0
		}
		if len(playerHand) == 1 {
			playerHand = append(playerHand, drawCard(deck))
			fmt.Printf("Your split hand is dealt a %v.\n", playerHand[1])
		}
		fmt.Printf("Your hand: %v, with a total value of %v.\n", playerHand, getHandValue(playerHand))
		fmt.Printf("The house has an %v (value of %v), and another card face down.\n", houseHand[0], getHandValue([]Card{houseHand[0]}))
		fmt.Println("What will you do?")
		fmt.Println("1 - Hit")
		fmt.Println("2 - Stand")
		fmt.Println("3 - Double Down")
		if len(playerHand) == 2 && playerHand[0].face == playerHand[1].face { // if player has option to split (assumes playercard1&2 exist)
			fmt.Println("4 - Split")
		}

		fmt.Scan(&gameInput)
		switch gameInput {
		case 1:
			// Hit
			playerHand = append(playerHand, drawCard(deck))
		case 2:
			// Stand
			clearCLI()
			fmt.Printf("The house reveals their face down card to be a %v.\n", houseHand[1])
			houseTurn(deck, playerHand, houseHand, bet)
			return 0
		case 3:
			// Double down
			bet *= 2
		case 4:
			if !(playerHand[0].face == playerHand[1].face) {
				fmt.Println("Please enter a valid input.")
				fmt.Printf("%v, %v, %v\n", gameInput, playerHand[0].face, playerHand[1].face)
			} else {
				// Split
				playerTurn(deck, []Card{playerHand[0]}, houseHand, bet)
				playerTurn(deck, []Card{playerHand[1]}, houseHand, bet)
				return 0
			}
		}
	}
	return 0
}

func houseTurn(deck []Card, playerHand []Card, houseHand []Card, bet int) {
	var houseScreenInput int
	for {
		fmt.Printf("The house hand is: %v, with a total value of %v.\n", houseHand, getHandValue(houseHand))
		if getHandValue(houseHand) > 21 {
			fmt.Println("The house has gone bust!")
			fmt.Println("Enter 1 to continue...")
			fmt.Scan(&houseScreenInput)
			adjustBalance(1, bet)
			break
		}
		if getHandValue(houseHand) > getHandValue(playerHand) {
			fmt.Println("House wins.")
			fmt.Println("Enter 1 to continue...")
			fmt.Scan(&houseScreenInput)
			adjustBalance(2, bet)
			break
		}
		if getHandValue(houseHand) >= 17 {
			if getHandValue(houseHand) != 21 {
				fmt.Println("The house stands because it cannot hit past 17.")
			}
			if getHandValue(playerHand) > getHandValue(houseHand) {
				// Player wins
				fmt.Println("Player wins.")
				fmt.Println("Enter 1 to continue...")
				fmt.Scan(&houseScreenInput)
				adjustBalance(1, bet)
				break
			} else if getHandValue(playerHand) == getHandValue(houseHand) {
				// Tie
				fmt.Println("Tie.")
				fmt.Println("Enter 1 to continue...")
				fmt.Scan(&houseScreenInput)
				adjustBalance(3, bet)
				break
			}
		}
		fmt.Println("Enter 1 to continue...")
		fmt.Scan(&houseScreenInput)
		houseHand = append(houseHand, drawCard(deck))
		clearCLI()
		fmt.Printf("The house draws another card. It is a %v.\n", houseHand[len(houseHand)-1])
	}
}

func main() {
	var mainScreenInput int
	var username string

	fmt.Println("Welcome to the Blackjack Casino! Please enter your name.")
	fmt.Scan(&username)
	// validate username
	clearCLI()

	fmt.Printf("Welcome %v. Please input a number according to the options below:\n", username)
	startScreen()

	for mainScreenInput != 3 {
		clearCLI()
		playerHand := []Card{}
		houseHand := []Card{}
		deck = initDeck()
		playerHand = append(playerHand, drawCard(deck), drawCard(deck))
		houseHand = append(houseHand, drawCard(deck), drawCard(deck))
		
		var bet int = placeBet()
		playerTurn(deck, playerHand, houseHand, bet)
		clearCLI()
		fmt.Println("Thanks for playing.")
		fmt.Println("1 - Play another round")
		fmt.Println("2 - Exit to menu")
		fmt.Scan(&mainScreenInput)
		if mainScreenInput == 1 {
			continue
		} else if mainScreenInput == 2 {
			clearCLI()
			startScreen()
		}
	}
	fmt.Println("Stop")
	ledgerScreen()
}

func startScreen() {
	var startScreenInput uint
	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Balance")
	fmt.Println("4 - Ledger")
	fmt.Println("5 - Quit game")
	fmt.Scan(&startScreenInput)
	switch startScreenInput {
	case 2:
		rulesScreen()
	case 3:
		balanceScreen()
	case 4:
		ledgerScreen()
	case 5:
		os.Exit(0)
	}
}

func rulesScreen() {
	var menuScreenInput int
	clearCLI()

	fmt.Println("******************** RULES ********************")
	fmt.Println("In Blackjack, you play against a dealer.")
	fmt.Println("The goal of the game is to build a hand with a total value as close to 21 as possible, without going over.")
	fmt.Println("The value of a card is its face value, with the exception of face cards which always have a value of 10, and aces whose value can either be 1 or 11.")
	fmt.Println("First, a player and dealer both draw a card.")
	fmt.Println("The player and house must then decide whether to 'hit' (draw again) or 'stay' (keep current total).")
	fmt.Println("The game ends when both the house and player stop drawing cards, or when one party has a total over 21.")
	fmt.Println("Other rules like the 17 pickup rule for the dealer are not *yet* implemented. This is basic Blackjack for now.")
	fmt.Println("\n1 - Play")
	fmt.Println("2 - Exit to menu")
	fmt.Println("***********************************************")
	fmt.Scan(&menuScreenInput)
	clearCLI()

	if menuScreenInput == 1 {
		// play
	} else if menuScreenInput == 2 {
		startScreen()
	}
}

func balanceScreen() {
	var balanceScreenInput int
	var addedFundsAmount int
	for balanceScreenInput != 2	{
		clearCLI()
		// balance
		fmt.Printf("You have a current balance of $%v.\n", balance)
		fmt.Println("1 - Add to balance")
		fmt.Println("2 - Exit to menu")
		fmt.Scan(&balanceScreenInput)
		if balanceScreenInput == 1 {
			fmt.Println("Enter an amount to add to your balance:")
			fmt.Scan(&addedFundsAmount)
			balance += addedFundsAmount
		}
	}
	clearCLI()
	startScreen()
}

func ledgerScreen() {
	var ledgerScreenInput int
	clearCLI()
	fmt.Println("The following is an ordered log of placed bets:")
	fmt.Println(ledger)
	fmt.Println("Enter 1 to continue...")
	fmt.Scan(&ledgerScreenInput)
	clearCLI()
	startScreen()
	// ledger
}