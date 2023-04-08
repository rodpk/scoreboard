package handler

import (
	"github.com/google/uuid"
	"github.com/rodpk/scoreboard/internal/model"
	"github.com/rodpk/scoreboard/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScoreboardServiceImpl struct {
	repository repository.ScoreboardRepository
}

func NewScoreboardService(r repository.ScoreboardRepository) ScoreboardHandler {
	return &ScoreboardServiceImpl{
		repository: r,
	}
}

func (*ScoreboardServiceImpl) CreateScoreboard(*model.Scoreboard) error {
	panic("unimplemented")
}

func (ss *ScoreboardServiceImpl) ListScoreboards(sbFilters *model.ScoreboardFilters) ([]model.Scoreboard, error) {

	filters := bson.D{}
	if sbFilters.Title != "" {
		filters = append(filters, primitive.E{Key: "Title", Value: sbFilters.Title})
	}

	if sbFilters.PlayerName != "" {
		filters = append(filters, primitive.E{
			Key: "Players",
			Value: bson.M{
				"$elemMatch": bson.M{
					"Name": primitive.Regex{Pattern: sbFilters.PlayerName, Options: "i"},
				},
			},
		})
	}

	return ss.repository.ListScoreboards(filters)
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
