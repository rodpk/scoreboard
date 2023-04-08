package utils

import (
	"github.com/google/uuid"
	"github.com/rodpk/scoreboard/internal/model"
)

// FindPlayerInScoreboard returns the index of the found player, returns -1 if not found
func FindPlayerInScoreboard(playerID uuid.UUID, players []model.Player) int {
	index := -1
	
	for i, p := range players {
		if p.Id == playerID {
			index = i
		}
	}

	return index
}