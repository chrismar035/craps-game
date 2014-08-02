package main

import (
	"github.com/chrismar035/go-craps"
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Welcome to Craps!")
	showMenu()

	game := craps.CrapsGame{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choice := scanner.Text()

		switch choice {
		case "q":
			fmt.Println("Thanks for playing!")
			return
		case "r":
			game.Roll()
			fmt.Printf("Roll: %v, %v - %v\n",
					   game.FirstDie(),
					   game.SecondDie(),
					   game.RollValue())
		case "b":
			amount := 10
			fmt.Printf("Placing $%v on the pass line\n", amount)
			game.PlaceBet(craps.PlayerBet{amount, craps.PassLineOdds})
			fmt.Printf("Bets: %v\n", len(game.Bets))
		}

		if game.IsComingOut() {
			fmt.Println("Coming out!")
		} else {
			fmt.Printf("Point is %v\n", game.Point)
		}

		showWinners(game.Winners)
		showBets(game.Bets)

		showMenu()
	}
}

func showMenu() {
	fmt.Println("\n\nMake your next move:")
	fmt.Println("r - Roll the dice!")
	fmt.Println("b - Place a bet")
	fmt.Println("q - Quit\n")
}


func showBets(bets []craps.PlayerBet) {
	if len(bets) == 0 {
		fmt.Println("No bets on the table")
	} else {
		fmt.Println("Current bets are:")
		for _, bet := range bets {
			fmt.Printf("%v at %v:%v\n", bet.Amount, bet.ToPay(), bet.Base())
		}
	}
}

func showWinners(winners []craps.PlayerBet) {
	if len(winners) == 0 {
		return
	}

	fmt.Println("Congrats to our winners:")
	for _, win := range winners {
		fmt.Printf("%v\n", win.Winnings())
	}

}
