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

var deck []Card
func initDeck() []Card {
	deck := []Card{
	{face: "Ace", suit: "Spades"},
	{face: "Ace", suit: "Hearts"},
	{face: "Ace", suit: "Diamonds"},
	{face: "Ace", suit: "Clubs"},
	{face: "Two", suit: "Spades", value: 2},
	{face: "Two", suit: "Hearts", value: 2},
	{face: "Two", suit: "Diamonds", value: 2},
	{face: "Two", suit: "Clubs", value: 2},
	{face: "Three", suit: "Spades", value: 3},
	{face: "Three", suit: "Hearts", value: 3},
	{face: "Three", suit: "Diamonds", value: 3},
	{face: "Three", suit: "Clubs", value: 3},
	{face: "Four", suit: "Spades", value: 4},
	{face: "Four", suit: "Hearts", value: 4},
	{face: "Four", suit: "Diamonds", value: 4},
	{face: "Four", suit: "Clubs", value: 4},
	{face: "Five", suit: "Spades", value: 5},
	{face: "Five", suit: "Hearts", value: 5},
	{face: "Five", suit: "Diamonds", value: 5},
	{face: "Five", suit: "Clubs", value: 5},
	{face: "Six", suit: "Spades", value: 6},
	{face: "Six", suit: "Hearts", value: 6},
	{face: "Six", suit: "Diamonds", value: 6},
	{face: "Six", suit: "Clubs", value: 6},
	{face: "Seven", suit: "Spades", value: 7},
	{face: "Seven", suit: "Hearts", value: 7},
	{face: "Seven", suit: "Diamonds", value: 7},
	{face: "Seven", suit: "Clubs", value: 7},
	{face: "Eight", suit: "Spades", value: 8},
	{face: "Eight", suit: "Hearts", value: 8},
	{face: "Eight", suit: "Diamonds", value: 8},
	{face: "Eight", suit: "Clubs", value: 8},
	{face: "Nine", suit: "Spades", value: 9},
	{face: "Nine", suit: "Hearts", value: 9},
	{face: "Nine", suit: "Diamonds", value: 9},
	{face: "Nine", suit: "Clubs", value: 9},
	{face: "Ten", suit: "Spades", value: 10},
	{face: "Ten", suit: "Hearts", value: 10},
	{face: "Ten", suit: "Diamonds", value: 10},
	{face: "Ten", suit: "Clubs", value: 10},
	{face: "Jack", suit: "Spades", value: 10},
	{face: "Jack", suit: "Hearts", value: 10},
	{face: "Jack", suit: "Diamonds", value: 10},
	{face: "Jack", suit: "Clubs", value: 10},
	{face: "Queen", suit: "Spades", value: 10},
	{face: "Queen", suit: "Hearts", value: 10},
	{face: "Queen", suit: "Diamonds", value: 10},
	{face: "Queen", suit: "Clubs", value: 10},
	{face: "King", suit: "Spades", value: 10},
	{face: "King", suit: "Hearts", value: 10},
	{face: "King", suit: "Diamonds", value: 10},
	{face: "King", suit: "Clubs", value: 10}}
	return deck
}

func drawCard(deck []Card) Card {
	index := rand.Intn(len(deck))
	card := deck[index]
	deck = append(deck[:index], deck[index+1:]...)
	return card
}

func main() {
	var username string

	fmt.Println("Welcome to the Blackjack Casino! Please enter your name.")
	fmt.Scan(&username)
	// validate username

	fmt.Printf("Welcome %v. Please input a number according to the options below:\n", username)
	startScreen()

	playerHand := []Card{}
	houseHand := []Card{}
	deck = initDeck()
	fmt.Println("A new round of Blackjack has started.")
	// game loop
	for {
		var gameInput int
		//TODO testing some stuff for split below
		//var playerCard1 = drawCard(deck)
		//var playerCard2 = drawCard(deck)

		var playerCard1 = Card{face: "Ace", suit: "Spades"}
		var playerCard2 = Card{face: "Ace", suit: "Diamonds"}
		var houseCard1 = drawCard(deck)
		var houseCard2 = drawCard(deck)

		playerHand = append(playerHand, playerCard1, playerCard2)
		houseHand = append(houseHand, houseCard1, houseCard2)

		var hasOptionToSplit bool = gameInput == 0 && playerCard1.face == playerCard2.face

		playerTurn:
		for {
			if getHandValue(playerHand) > 21 {
				fmt.Println("Your hand went bust!")
				break
			}
			fmt.Printf("Your hand: %v, with a total value of %v.\n", playerHand, getHandValue(playerHand))
			fmt.Printf("The house has an %v (value of %v), and another card face down.\n", houseCard1, getHandValue([]Card{houseHand[0]}))
			fmt.Println("What will you do?")
			fmt.Println("1 - Hit")
			fmt.Println("2 - Stand")
			fmt.Println("3 - Double Down")
			if hasOptionToSplit {
				fmt.Println("4 - Split")
			}

			fmt.Scan(&gameInput)
			switch gameInput {
			case 1:
				// Hit
				playerHand = append(playerHand, drawCard(deck))
			case 2:
				// Stand
				break playerTurn
			case 3:
				// Double down
			case 4:
				if !hasOptionToSplit {
					fmt.Println("Please enter a valid input.")
				} else {
					// Split
					fmt.Println("test")
				}
			}

			fmt.Printf("%v, %v\n", playerHand, getHandValue(playerHand))
		}
		
		fmt.Printf("The house reveals their face down card to be a %v.\n", houseHand[1])
		houseTurn:
		for {
			if getHandValue(houseHand) > 21 {
				fmt.Println("The house has gone bust!")
				break
			}
			fmt.Printf("The house hand is: %v, with a total value of %v.\n", houseHand, getHandValue(houseHand))
			if getHandValue(houseHand) >= 17 {
				fmt.Printf("The house stands.")
				if getHandValue(playerHand) > getHandValue(houseHand) {
					// Player wins
					fmt.Println("Player wins")
					break houseTurn
				} else if getHandValue(playerHand) == getHandValue(houseHand) {
					// Tie
					fmt.Println("Tie")
					break houseTurn
				} else {
					// House wins
					fmt.Println("House wins")
					break houseTurn
				}
			}
			houseHand = append(houseHand, drawCard(deck))
			fmt.Printf("The house draws another card. It is a %v.", houseHand[len(houseHand)-1])
		}
		// start house turn
		//fmt.Printf("deck2%v", deck)
		var gameInputtmp int
		fmt.Scan(&gameInputtmp)
		if gameInputtmp == 2 {
			break
		} else if gameInputtmp == 1 {
			continue
		}

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
	var menuScreenInput int

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

	if menuScreenInput == 1 {
		// play
	} else if menuScreenInput == 2 {
		startScreen()
	}
}

func balanceScreen() {
	// balance
}

func ledgerScreen() {
	// ledger
}