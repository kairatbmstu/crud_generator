package repository

import (
	"database/sql"
	"example.com/ast1/test/entity"
	"github.com/google/uuid"
)

type StudentRepository struct {
	db *sql.DB
}

func (r *StudentRepository) Create(student *entity.Student) error {
	_, err := r.db.Exec("INSERT INTO students (id,firstname,lastname,age) VALUES ($1,$2,$3,$4)", student.Id, student.Firstname, student.Lastname, student.Age)
	return err
}
func (r *StudentRepository) Update(student *entity.Student) error {
	_, err := r.db.Exec("UPDATE  students set firstname = $2,lastname = $3,age = $4, WHERE id = $1", student.Id, student.Firstname, student.Lastname, student.Age)
	return err
}
func (r *StudentRepository) Delete(student *entity.Student) error {
	_, err := r.db.Exec("DELETE students  WHERE id = $1", student.Id)
	return err
}
func (r *StudentRepository) FindByID(id uuid.UUID) (*entity.Student, error) {
	_, err := r.db.QueryRow("SELECT id,firstname,lastname,age from students  WHERE id = $1", student.Id), &Student{Name: s.Name, Age: s.Age}
	if err != nil {
		return nil, err
	}
	return nil, err
}
