package game

import (
	"card-game/internal/deck"
	"card-game/internal/models"
	"sort"
	"strconv"
)

func deal() {
	emptyDeck := false
	for emptyDeck == false {
		for i := range table.Players {

			var (
				c *models.Card
				err error
			)
			c, table.Deck, err = deck.Draw(table.Deck)
			if err != nil {
				panic(err)
			}
			table.Players[i].DrawPile = append(table.Players[i].DrawPile, *c)

			if len(table.Deck) == 0 {
				emptyDeck = true
				break
			}
		}
	}
}

func players(names []string) []models.Player {
	players := make([]models.Player, len(names))
	for i, name := range names {
		players[i] = models.Player{
			Name:        name,
			DrawPile:    make([]models.Card, 0, 40),
			DiscardPile: make([]models.Card, 0, 40),
		}
	}

	return players
}

func win(winner *models.Player, info *string) (bool, *string, *string) {
	winner.DiscardPile = append(winner.DiscardPile, table.Unclaimed...)
	table.Unclaimed = make([]models.Card, 0, 40)
	for i := range table.Players {
		winner.DiscardPile = append(winner.DiscardPile, *table.Players[i].Hand)
		table.Players[i].Hand = nil
	}
	return true, &winner.Name, info
}

func even(info *string) (bool, *string, *string) {
	for i := range table.Players {
		table.Unclaimed = append(table.Unclaimed, *table.Players[i].Hand)
		table.Players[i].Hand = nil
	}
	return true, nil, info
}

func checkDeck(i int) {
	if len(table.Players[i].DrawPile) == 0 {
		if len(table.Players[i].DiscardPile) != 0 {
			deck.Shuffle(table.Players[i].DiscardPile)
			table.Players[i].DrawPile = table.Players[i].DiscardPile
			table.Players[i].DiscardPile = make([]models.Card, 0, 40)
		}
	}
}

func checkPlayer(i int) bool {
	if len(table.Players[i].DrawPile) == 0 && len(table.Players[i].DiscardPile) == 0 {
		table.Players[i] = table.Players[len(table.Players)-1]
		table.Players = table.Players[:len(table.Players)-1]
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

func hands(players []models.Player) *string {
	info := ""
	for _, player := range players {
		info += hand(player)
	}

	return &info
}

func hand(player models.Player) string {
	return player.Name + ": " + strconv.Itoa(player.Hand.Rank) + " " + player.Hand.Suit.String() + "; "
}
