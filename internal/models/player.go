package models

type Player struct {
	Name        string
	DrawPile    []Card
	DiscardPile []Card
	Hand *Card
}
