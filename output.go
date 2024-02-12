package main

import "database/sql"

type Student struct {
	ID   int
	Name string
	Age  int
}
type StudentRepository struct {
	db sql.DB
}

func Create(s *Student) error
func Update(s *Student) error
func Delete() error
func FindByID(idOrName string) (error, *Student)
func FindByName(idOrName string) (error, *Student)
func FindByAge(age int) (error, []*Student)
