package models

type Table struct {
	Players   []Player
	Deck      []Card
	Unclaimed []Card
}
