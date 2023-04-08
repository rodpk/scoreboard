package repository

import (
	"github.com/google/uuid"
	"github.com/rodpk/scoreboard/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScoreboardRepository interface {
	// CreateScoreboard creates a new scoreboard with the given details.
	// Returns an error if the operation fails.
	CreateScoreboard(*model.Scoreboard) error

	// ListScoreboards returns a list of all scoreboards that match the given filter.
	// Returns an error if the operation fails.
	ListScoreboards(filter primitive.D) ([]model.Scoreboard, error)

	// FindScoreboard finds and returns the scoreboard that matches the given filter.
	// Returns nil if no matching scoreboard is found.
	// Returns an error if the operation fails.
	FindScoreboard(filter primitive.D) (*model.Scoreboard, error)

	// AddPlayerToScoreboard adds a player to the specified scoreboard.
	// Returns the updated scoreboard and an error if the operation fails.
	AddPlayerToScoreboard(scoreboardID uuid.UUID, player model.Player) (*model.Scoreboard, error)

	// UpdatePlayerScore updates the score of a player in the specified scoreboard.
	// Returns the updated scoreboard and an error if the operation fails.
	UpdatePlayerScore(scoreboardID uuid.UUID, playerID uuid.UUID, value int) (*model.Scoreboard, error)

	// RemovePlayerFromScoreboard removes a player from the specified scoreboard.
	// Returns the updated scoreboard and an error if the operation fails.
	RemovePlayerFromScoreboard(scoreboardID uuid.UUID, playerID uuid.UUID) (*model.Scoreboard, error)

	// UpdateScoreboardTitle updates the title of the specified scoreboard.
	// Returns the updated scoreboard and an error if the operation fails.
	UpdateScoreboardTitle(scoreboardID uuid.UUID, newTitle string) (*model.Scoreboard, error)

	// CreateHistoryEntry creates a new change history entry with the given details.
	// Returns an error if the operation fails.
	CreateHistoryEntry(*model.ChangeHistory) error
}
