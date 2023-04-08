package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodpk/scoreboard/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScoreboardRepositoryImpl struct {
	client *mongo.Client
}

func NewScoreboardRepository(c *mongo.Client) ScoreboardRepository {
	return &ScoreboardRepositoryImpl{
		client: c,
	}
}

func (sr *ScoreboardRepositoryImpl) CreateScoreboard(s *model.Scoreboard) error {
	collection := sr.client.Database("scoreboard").Collection("scoreboards")
	_, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		return err
	}

	return nil
}

func (sr *ScoreboardRepositoryImpl) ListScoreboards(filter primitive.D) ([]model.Scoreboard, error) {
	collection := sr.client.Database("scoreboard").Collection("scoreboards")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	scoreboards := []model.Scoreboard{}
	for cur.Next(context.Background()) {
		var scoreboard model.Scoreboard
		err := cur.Decode(&scoreboard)
		if err != nil {
			return nil, err
		}
		scoreboards = append(scoreboards, scoreboard)
	}

	return scoreboards, nil
}
func (sr *ScoreboardRepositoryImpl) FindScoreboard(filter primitive.D) (*model.Scoreboard, error) {
	collection := sr.client.Database("scoreboard").Collection("scoreboards")
	var scoreboard model.Scoreboard
	err := collection.FindOne(context.Background(), filter).Decode(&scoreboard)
	if err != nil {
		return nil, err
	}
	return &scoreboard, nil
}

func (*ScoreboardRepositoryImpl) AddPlayerToScoreboard(scoreboardID uuid.UUID, player model.Player) (*model.Scoreboard, error) {
	panic("unimplemented")
}

func (*ScoreboardRepositoryImpl) CreateHistoryEntry(*model.ChangeHistory) error {
	panic("unimplemented")
}

func (*ScoreboardRepositoryImpl) RemovePlayerFromScoreboard(scoreboardID uuid.UUID, playerID uuid.UUID) (*model.Scoreboard, error) {
	panic("unimplemented")
}

func (*ScoreboardRepositoryImpl) UpdatePlayerScore(scoreboardID uuid.UUID, playerID uuid.UUID, value int) (*model.Scoreboard, error) {
	panic("unimplemented")
}

func (*ScoreboardRepositoryImpl) UpdateScoreboardTitle(scoreboardID uuid.UUID, newTitle string) (*model.Scoreboard, error) {
	panic("unimplemented")
}
