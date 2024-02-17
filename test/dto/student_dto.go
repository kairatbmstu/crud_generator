package model

import "github.com/google/uuid"

type StudentDTO struct {
	Id        uuid.UUID
	Firstname string
	Lastname  string
	Age       int
}
