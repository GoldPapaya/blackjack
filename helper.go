package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

// Card object
type Card struct {
	face string
	suit string
	value int
}

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

func placeBet() int {
	var betScreenInput int
	var bet int
	fmt.Printf("Your current balance is $%v.\n", balance)
	fmt.Println("Enter the amount of your bet:")
	fmt.Scan(&bet)
	fmt.Printf("Your bet of $%v has been placed.\n", bet)
	fmt.Println("Enter 1 to continue...")
	fmt.Scan(&betScreenInput)
	return bet
}

func adjustBalance(state int, bet int) {
	clearCLI()
	var adjustBalanceScreenInput int
	switch state {
	case 1:
		balance += bet
		ledger = append(ledger, bet)
		fmt.Printf("You won $%v!\n", bet)
	case 2:
		balance -= bet
		ledger = append(ledger, -bet)
		fmt.Printf("You lost $%v!\n", bet)
	case 3:
		ledger = append(ledger, 0)
		fmt.Println("You tied!")
	}
	fmt.Printf("Your balance is now $%v.\n", balance)
	fmt.Println("Enter 1 to continue...")
	fmt.Scan(&adjustBalanceScreenInput)
}

func drawCard(deck []Card) Card {
	index := rand.Intn(len(deck))
	card := deck[index]
	deck = append(deck[:index], deck[index+1:]...)
	return card
}

// For clearing CLI output only
func clearCLI() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else { // this might be problematic for a non-windows os
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
