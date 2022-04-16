package main

import (
	"card-game/pkg/cardgame"
	"fmt"
	"time"
)

const green = "\033[32m"
const red = "\033[31m"
const reset = "\033[0m"
const separator = "---------------------------------------------------------------"

var (
	game     cardgame.Game
	deckSize = 400
	players  = []string{"Bogdan", "Alex", "Pera", "Zika"}
)

func main() {
	game = cardgame.Game{}
	t := time.Now()
	defer gameTime(t)

	err := game.Setup(deckSize, players)

	if err != nil {
		fmt.Print(red, fmt.Sprintf("Setup incorrect: %s\n", err))
	}

	fmt.Println("Game started")
	fmt.Println(separator)
	for running, err := game.Play(); running; running, err = game.Play() {
		if err != nil {
			fmt.Print(red, fmt.Sprintf("%s\n", err))
			break
		}

		fmt.Println(game.RoundInfo())
		if game.RoundWinner() == nil {
			fmt.Println("Draw")
			continue
		}

		fmt.Printf("Round winner is %s\n", *game.RoundWinner())
		fmt.Println(separator)
	}
	fmt.Println(separator)
	fmt.Print(green, fmt.Sprintf("Game winner is %s\n", game.Winner()))
}

func gameTime(start time.Time) {
	fmt.Print(reset, fmt.Sprintf("Game lasted %s\n", time.Since(start)))
}
