package model

import "github.com/google/uuid"

type Student  struct {
	Id   uuid.UUID
	Name string
	Age  int
}
