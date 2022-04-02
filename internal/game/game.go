package game

import (
	"card-game/internal/deck"
	"card-game/internal/models"
)

var (
	table     models.Table
	setupDone bool
	winner    string
)

func Setup(deckSize int, names ...string) {
	var err error
	table = models.Table{}
	table.Players = players(names)
	table.Deck, err = deck.Create(deckSize)

	if err != nil {
		panic(err.Error())
	}

	deck.Shuffle(table.Deck)
	deal()
	setupDone = true
}

func Play() (bool, *string, *string) {
	if setupDone == false {
		panic("Setup not done")
	}

	for i := 0; i < len(table.Players); i++ {
		if len(table.Players) == 1 {
			winner = table.Players[0].Name
			return false, &table.Players[0].Name, nil
		}
		checkDeck(i)
		if !checkPlayer(i) {
			i--
			continue
		}

		var err error
		table.Players[i].Hand, table.Players[i].DrawPile, err = deck.Draw(table.Players[i].DrawPile)

		if err != nil {
			panic(err)
		}
	}

	if w := evaluate(table.Players); w == nil {
		return even(hands(table.Players))
	} else {
		return win(w, hands(table.Players))
	}
}

func Winner() string {
	return winner
}
