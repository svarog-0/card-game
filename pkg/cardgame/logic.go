package cardgame

import (
	"card-game/internal/deck"
	"card-game/internal/models"
	"sort"
	"strconv"
)

func (game *Game) deal() {
	emptyDeck := false
	for !emptyDeck {
		for i := range game.table.Players {

			var (
				c   *models.Card
				err error
			)
			c, game.table.Deck, err = deck.Draw(game.table.Deck)
			if err != nil {
				panic(err)
			}
			game.table.Players[i].DrawPile = append(game.table.Players[i].DrawPile, *c)

			if len(game.table.Deck) == 0 {
				emptyDeck = true
				break
			}
		}
	}
}

func (game *Game) players(names []string) []models.Player {
	players := make([]models.Player, len(names))
	for i, name := range names {
		players[i] = models.Player{
			Name:        name,
			DrawPile:    make([]models.Card, 0, game.deckSize),
			DiscardPile: make([]models.Card, 0, game.deckSize),
		}
	}

	return players
}

func (game *Game) win(winner *models.Player) *string {
	winner.DiscardPile = append(winner.DiscardPile, game.table.Unclaimed...)
	game.table.Unclaimed = make([]models.Card, 0, game.deckSize)
	for i := range game.table.Players {
		winner.DiscardPile = append(winner.DiscardPile, *game.table.Players[i].Hand)
		game.table.Players[i].Hand = nil
	}

	return &winner.Name
}

func (game *Game) even() *string {
	for i := range game.table.Players {
		game.table.Unclaimed = append(game.table.Unclaimed, *game.table.Players[i].Hand)
		game.table.Players[i].Hand = nil
	}

	return nil
}

func (game *Game) checkDeck(i int) {
	if len(game.table.Players[i].DrawPile) == 0 {
		if len(game.table.Players[i].DiscardPile) != 0 {
			deck.Shuffle(game.table.Players[i].DiscardPile)
			game.table.Players[i].DrawPile = game.table.Players[i].DiscardPile
			game.table.Players[i].DiscardPile = make([]models.Card, 0, game.deckSize)
		}
	}
}

func (game *Game) checkPlayer(i int) bool {
	if len(game.table.Players[i].DrawPile) == 0 && len(game.table.Players[i].DiscardPile) == 0 {
		game.table.Players[i] = game.table.Players[len(game.table.Players)-1]
		game.table.Players = game.table.Players[:len(game.table.Players)-1]
		return false
	}

	return true
}

func evaluate(players []models.Player) *models.Player {
	if len(players) == 1 {
		return &players[0]
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Hand.Rank > players[j].Hand.Rank
	})

	if players[0].Hand.Rank == players[1].Hand.Rank {
		return nil
	}

	return &players[0]
}

func hands(players []models.Player) string {
	info := ""
	for _, player := range players {
		info += hand(player)
	}

	return info
}

func hand(player models.Player) string {
	return player.Name + ": " + strconv.Itoa(player.Hand.Rank) + " " + player.Hand.Suit.String() + "; "
}
