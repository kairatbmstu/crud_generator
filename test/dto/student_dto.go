package model

import "github.com/google/uuid"

type StudentDTO struct {
	Id   uuid.UUID
	Name string
	Age  int
}
