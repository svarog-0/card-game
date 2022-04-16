package cardgame

import (
	"card-game/internal/deck"
	"card-game/internal/models"
	"errors"
)

type Game struct {
	deckSize    int
	names       []string
	table       models.Table
	setupDone   bool
	winner      string
	roundInfo   string
	roundWinner *string
}

func (game *Game) Setup(deckSize int, names []string) error {
	var err error
	game.deckSize = deckSize
	game.names = names
	game.table = models.Table{}
	game.table.Players = game.players(game.names)
	game.table.Deck, err = deck.Create(game.deckSize)

	if err != nil {
		return err
	}

	deck.Shuffle(game.table.Deck)
	game.deal()
	game.setupDone = true

	return nil
}

func (game *Game) Play() (bool, error) {
	if !game.setupDone {
		return false, errors.New("Setup not done")
	}

	for i := 0; i < len(game.table.Players); i++ {
		if len(game.table.Players) == 1 {
			game.roundWinner = nil
			game.roundInfo = ""
			game.winner = game.table.Players[0].Name
			return false, nil
		}

		game.checkDeck(i) /// change logic to be more nicer, all players play moves then deck&player check routine then all players play moves etc.
		if !game.checkPlayer(i) {
			i--
			game.roundWinner = nil
			game.roundInfo = ""
			continue
		}

		var err error
		game.table.Players[i].Hand, game.table.Players[i].DrawPile, err = deck.Draw(game.table.Players[i].DrawPile)

		if err != nil {
			panic(err)
		}
	}

	game.roundInfo = hands(game.table.Players)

	if winner := evaluate(game.table.Players); winner == nil {
		game.roundWinner = game.even()
	} else {
		game.roundWinner = game.win(winner)
	}

	return true, nil
}

func (game *Game) Winner() string {
	return game.winner
}

func (game *Game) RoundWinner() *string {
	return game.roundWinner
}

func (game *Game) RoundInfo() string {
	return game.roundInfo
}
