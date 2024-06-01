package gamesrv
import "v1/internal/core/domain"

func (game *Game) AddPlayer(player Player) {
	game.Players = append(game.Players, player)
}

func (game *Game) GetPlayer(playerID uint) (*Player, error) {
	for i, player := range game.Players {
		if player.ID == playerID {
			return &game.Players[i], nil
		}
	}
	return nil, errors.New("player not found")
}

func (game *Game) RemovePlayer(playerID uint) error {
	for i, player := range game.Players {
		if player.ID == playerID {
			game.Players = append(game.Players[:i], game.Players[i+1:]...)
			return nil
		}
	}
	return errors.New("player not found")
}
