package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rodpk/scoreboard/internal/model"
	"github.com/rodpk/scoreboard/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
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
	return err
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

func (sr *ScoreboardRepositoryImpl) AddPlayerToScoreboard(scoreboardID uuid.UUID, player model.Player) (*model.Scoreboard, error) {
	collection := sr.client.Database("scoreboard").Collection("scoreboards")
	filter := bson.D{{Key: "Id", Value: scoreboardID}}
	scoreboard, err := sr.FindScoreboard(filter)
	if err != nil {
		return nil, err
	}
	scoreboard.Players = append(scoreboard.Players, player)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "players", Value: scoreboard.Players}}}}

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return nil, err
	}

	// log res
	fmt.Printf("res: %v\n", res)

	return scoreboard, nil
}

func (sr *ScoreboardRepositoryImpl) CreateHistoryEntry(ch *model.ChangeHistory) error {
	collection := sr.client.Database("scoreboard").Collection("histories")
	_, err := collection.InsertOne(context.Background(), &ch)
	return err
}

func (sr *ScoreboardRepositoryImpl) RemovePlayerFromScoreboard(scoreboardID uuid.UUID, playerID uuid.UUID) (*model.Scoreboard, error) {
	collection := sr.client.Database("scoreboard").Collection("scoreboards")
	filter := bson.D{{Key: "Id", Value: scoreboardID}}
	scoreboard, err := sr.FindScoreboard(filter)

	if err != nil {
		return nil, err
	}

	index := utils.FindPlayerInScoreboard(playerID, scoreboard.Players)

	if index == -1 {
		return nil, errors.New("player not found in scoreboard")
	}

	// remove player from slice
	scoreboard.Players = append(scoreboard.Players[:index], scoreboard.Players[index+1:]...)

	// update
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "Players", Value: scoreboard.Players}}}}
	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return nil, err
	}

	// todo
	fmt.Printf("res: %v\n", res)
	return scoreboard, nil
}

func (sr *ScoreboardRepositoryImpl) UpdatePlayerScore(scoreboardID uuid.UUID, playerID uuid.UUID, value int) (*model.Scoreboard, error) {
	collection := sr.client.Database("scoreboard").Collection("scoreboards")
	scoreboard, err := sr.FindScoreboard(bson.D{{Key: "Id", Value: scoreboardID}})
	if err != nil {
		return nil, err
	}

	// find player with specified ID within the scoreboard
	index := utils.FindPlayerInScoreboard(playerID, scoreboard.Players)
	if index == -1 {
		return nil, fmt.Errorf("player not found in scoreboard")
	}

	// update player's score
	scoreboard.Players[index].Score = value

	// update scoreboard in database
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "Players", Value: scoreboard.Players}}}}
	filter := bson.D{{Key: "Id", Value: scoreboardID}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return scoreboard, nil
}

func (sr *ScoreboardRepositoryImpl) UpdateScoreboardTitle(scoreboardID uuid.UUID, title string) (*model.Scoreboard, error) {
	collection := sr.client.Database("scoreboard").Collection("scoreboards")
	filter := bson.D{{Key: "Id", Value: scoreboardID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "Title", Value: title}}}}
	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return nil, err
	}

	// todo
	fmt.Printf("res: %v\n", res)

	return sr.FindScoreboard(filter)
}
