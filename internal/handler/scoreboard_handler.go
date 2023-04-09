package handler

import (
	"net/http"
)

type ScoreboardHandler interface {
	CreateScoreboard(w http.ResponseWriter, r *http.Request)
	ListScoreboards(w http.ResponseWriter, r *http.Request)
	FindScoreboard(w http.ResponseWriter, r *http.Request)
	AddPlayerToScoreboard(w http.ResponseWriter, r *http.Request)
	UpdatePlayerScore(w http.ResponseWriter, r *http.Request)
	RemovePlayerFromScoreboard(w http.ResponseWriter, r *http.Request)
	UpdateScoreboardTitle(w http.ResponseWriter, r *http.Request)
}
