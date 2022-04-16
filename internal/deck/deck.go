package deck

import (
	"card-game/internal/models"
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Create(size int) ([]models.Card, error) {
	if size < 4 {
		return nil, errors.New("minimum deck size is 4")
	}
	if l := size % 4; l != 0 {
		return nil, errors.New("deck size must be dividable by 4")
	}

	perSuite := size / 4
	deck := make([]models.Card, 0, size)
	deck = append(deck, cards(perSuite, models.Clubs)...)
	deck = append(deck, cards(perSuite, models.Hearts)...)
	deck = append(deck, cards(perSuite, models.Diamonds)...)
	deck = append(deck, cards(perSuite, models.Spades)...)

	return deck, nil
}

func Shuffle(list []models.Card) {
	n := len(list)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

func Draw(deck []models.Card) (*models.Card, []models.Card, error) {
	if len(deck) == 0 {
		return nil, nil, errors.New("deck is empty")
	}
	card := deck[0]
	if len(deck) == 1 {
		return &card, make([]models.Card, 0), nil
	}
	deck = deck[1:]
	return &card, deck, nil
}

func cards(size int, suit models.Suits) []models.Card {
	deck := make([]models.Card, size)
	for i := 0; i < size; i++ {
		deck[i] = models.Card{Rank: i + 1, Suit: suit}
	}

	return deck
}
