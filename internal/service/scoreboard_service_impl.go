package service

import (
	"github.com/google/uuid"
	"github.com/rodpk/scoreboard/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScoreboardServiceImpl struct{}

func NewScoreboardService( /*dependencies*/ ) ScoreboardService {
	return &ScoreboardServiceImpl{}
}

func (*ScoreboardServiceImpl) CreateScoreboard(*model.Scoreboard) error {
	panic("unimplemented")
}

func (*ScoreboardServiceImpl) ListScoreboards(filter primitive.D) ([]model.Scoreboard, error) {
	panic("unimplemented")
}
func (*ScoreboardServiceImpl) FindScoreboard(filter primitive.D) (*model.Scoreboard, error) {
	panic("unimplemented")
}

func (*ScoreboardServiceImpl) AddPlayerToScoreboard(scoreboardID uuid.UUID, player model.Player) (*model.Scoreboard, error) {
	panic("unimplemented")
}

func (*ScoreboardServiceImpl) CreateHistoryEntry(*model.ChangeHistory) error {
	panic("unimplemented")
}

func (*ScoreboardServiceImpl) RemovePlayerFromScoreboard(scoreboardID uuid.UUID, playerID uuid.UUID) (*model.Scoreboard, error) {
	panic("unimplemented")
}

func (*ScoreboardServiceImpl) UpdatePlayerScore(scoreboardID uuid.UUID, playerID uuid.UUID, value int) (*model.Scoreboard, error) {
	panic("unimplemented")
}

func (*ScoreboardServiceImpl) UpdateScoreboardTitle(scoreboardID uuid.UUID, newTitle string) (*model.Scoreboard, error) {
	panic("unimplemented")
}
