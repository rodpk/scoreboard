package model

import "github.com/google/uuid"


type Player struct {
	Id           uuid.UUID
	Name         string
	Score        int
}