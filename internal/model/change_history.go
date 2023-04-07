package model

import "github.com/google/uuid"



type ChangeHistory struct {
	Id            uuid.UUID
	ScoreboardId  uuid.UUID
	PlayerId      uuid.UUID
	PreviousScore int
	UpdatedScore  int
}