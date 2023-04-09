package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rodpk/scoreboard/internal/model"
	"github.com/rodpk/scoreboard/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseBody struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

type ScoreboardHandlerImpl struct {
	repository repository.ScoreboardRepository
}

func NewScoreboardHandler(r repository.ScoreboardRepository) ScoreboardHandler {
	return &ScoreboardHandlerImpl{
		repository: r,
	}
}

func (*ScoreboardHandlerImpl) AddPlayerToScoreboard(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (sh *ScoreboardHandlerImpl) CreateScoreboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var scoreboard model.Scoreboard
	err := json.NewDecoder(r.Body).Decode(&scoreboard)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseBody{
			Status:    http.StatusBadRequest,
			Data:      nil,
			Message:   fmt.Sprintf("Failed to decode request body: %v", err.Error()),
			Timestamp: time.Now(),
		})
		return
	}

	err = sh.repository.CreateScoreboard(&scoreboard)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseBody{
			Status:    http.StatusInternalServerError,
			Data:      nil,
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}
}

func (*ScoreboardHandlerImpl) FindScoreboard(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (sh *ScoreboardHandlerImpl) ListScoreboards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sbFilters := model.ScoreboardFilters{
		Title:      chi.URLParam(r, "title"),
		PlayerName: chi.URLParam(r, "playerName"),
	}

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

	scoreboards, err := sh.repository.ListScoreboards(filters)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseBody{
			Status:    http.StatusInternalServerError,
			Data:      nil,
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	responseBody := ResponseBody{
		Message: "Scoreboards retrieved successfully",
		Data:    scoreboards,
		Status:  http.StatusOK,
		Timestamp: time.Now(),
	}

	json.NewEncoder(w).Encode(responseBody)
}

func (*ScoreboardHandlerImpl) RemovePlayerFromScoreboard(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (*ScoreboardHandlerImpl) UpdatePlayerScore(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (*ScoreboardHandlerImpl) UpdateScoreboardTitle(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
