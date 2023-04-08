package model

import "github.com/google/uuid"

type Scoreboard struct {
	Id      uuid.UUID
	Title   string
	Players []Player
}
