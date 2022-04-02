package main

import (
	"card-game/internal/game"
	"fmt"
	"time"
)

const green = "\033[32m"
const red = "\033[31m"
const reset = "\033[0m"

func main() {
	defer gameTime(time.Now())
	defer failHandle()

	game.Setup(400, "Bogdan", "Alex", "Pera", "Zika")

	for running, winner, info := game.Play(); running; running, winner, info = game.Play() {
		fmt.Println(*info)
		if winner == nil {
			fmt.Println("Draw")
			continue
		}
		fmt.Printf("Round winner is %s\n", *winner)
	}
	fmt.Println()

	fmt.Print(green, fmt.Sprintf("Game winner is %s\n", game.Winner()))
}

func failHandle() {
	if r := recover(); r != nil {
		fmt.Print(red, fmt.Sprintf("ERROR: %s\n", r))
		fmt.Print(red, "Cannot continue game. Fatal fail!\n")
	}
}

func gameTime(start time.Time){
	fmt.Print(reset, fmt.Sprintf("Game lasted %s\n", time.Since(start)))
}
