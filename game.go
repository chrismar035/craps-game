package main

import (
	"bufio"
	"fmt"
	"github.com/chrismar035/go-craps"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Welcome to Craps!")
	showMenu()

	bets := []craps.PlayerBet{}
	point := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choice := scanner.Text()

		switch choice {
		case "q":
			fmt.Println("Thanks for playing!")
			return
		case "r":
			roll1 := craps.Roll()
			roll2 := craps.Roll()
			rollValue := roll1 + roll2
			fmt.Printf("Roll: %v, %v - %v\n", roll1, roll2, rollValue)
			if point == 0 {
				point = rollValue
			} else if rollValue == 7 {
				fmt.Println("Seven out! All bets away")
				bets = nil
				point = 0
			} else if rollValue == point {
				fmt.Println("Winner winner chicken dinner!")
				for _, bet := range bets {
					fmt.Printf("Player wins %v\n", bet.Winnings())
				}
				bets = nil
				point = 0
			}
		case "b":
			bets = append(bets, craps.PlayerBet{10, craps.PassLineOdds})
		}
		if point != 0 {
			fmt.Printf("Point is %v\n", point)
		} else {
			fmt.Println("Coming out!")
		}

		showBets(bets)

		showMenu()
	}
}

func showMenu() {
	fmt.Println("\n\nPlease choose one of the following options:")
	fmt.Println("r - Roll the dice!")
	fmt.Println("b - Place a bet")
	fmt.Println("q - Quit\n")
}

func showBets(bets []craps.PlayerBet) {
	fmt.Println("Current bets are:")
	for _, bet := range bets {
		fmt.Printf("%v at %v:%v\n", bet.Amount, bet.ToPay(), bet.Base())
	}
}
