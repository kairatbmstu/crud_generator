package dto

import "github.com/google/uuid"

type GroupDTO struct {
	Id        uuid.UUID
	Code      string
	StartYear int
}
