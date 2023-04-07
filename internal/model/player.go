package model

import "github.com/google/uuid"


type Players struct {
	Id           uuid.UUID
	Name         string
	Score        int
}