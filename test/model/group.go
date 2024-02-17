package model

import "github.com/google/uuid"

type Group struct {
	Id        uuid.UUID
	Code      string
	StartYear int
}
