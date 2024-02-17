package entity

import "github.com/google/uuid"

type Student struct {
	Id        uuid.UUID
	Firstname string
	Lastname  string
	Age       int
}
