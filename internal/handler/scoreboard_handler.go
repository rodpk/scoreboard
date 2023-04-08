package handler

import (
	"github.com/google/uuid"
	"github.com/rodpk/scoreboard/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScoreboardHandler interface {
	CreateScoreboard(*model.Scoreboard) error
	ListScoreboards(sbFilters *model.ScoreboardFilters) ([]model.Scoreboard, error)
	FindScoreboard(filter primitive.D) (*model.Scoreboard, error)
	AddPlayerToScoreboard(scoreboardID uuid.UUID, player model.Player) (*model.Scoreboard, error)
	UpdatePlayerScore(scoreboardID uuid.UUID, playerID uuid.UUID, value int) (*model.Scoreboard, error)
	RemovePlayerFromScoreboard(scoreboardID uuid.UUID, playerID uuid.UUID) (*model.Scoreboard, error)
	UpdateScoreboardTitle(scoreboardID uuid.UUID, newTitle string) (*model.Scoreboard, error)
	CreateHistoryEntry(*model.ChangeHistory) error
}
