package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rodpk/scoreboard/internal/handler"
)

func InitializeRoutes(r *chi.Mux, sh handler.ScoreboardHandler) {

	r.Route("/v1", func(r chi.Router) {
		r.Route("/scoreboards", func(r chi.Router) {
			r.Get("/", sh.ListScoreboards)
			r.Post("/", sh.CreateScoreboard)
		})

		r.Route("/health-check", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
		})
	})
}
