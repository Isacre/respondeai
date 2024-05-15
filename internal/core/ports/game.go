package ports 

import "v1/internal/core/domain"

type GamesRepository interface {
	AddPlayer(player domain.Player) 
	GetPlayer(playerId uint) (*domain.Player, error)
	RemovePlayer(playerId uint) (error)
}

