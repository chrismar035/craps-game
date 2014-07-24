package main

import (
	"github.com/chrismar035/go-craps"
	"fmt"
)

func main() {
	bet := craps.PlayerBet{10, craps.PassLineOdds}
	fmt.Printf("Winnings: %v\n", bet.Winnings())
}
