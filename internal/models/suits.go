package models

type Suits int64

const (
	Spades Suits = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suits) String() string {
	return [...]string{"♤ Spades", "♥ Hearts", "♢ Diamonds", "♧ Clubs"}[s]
}
