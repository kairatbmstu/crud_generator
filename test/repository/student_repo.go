package repository

import (
	"database/sql"
	"example.com/ast1/test/entity"
	"example.com/ast1/test/repository"
)

type StudentRepository struct {
	db *sql.DB
}

func (r *StudentRepository) Create(student *entity.Student) error {
	_, err := r.db.Exec("INSERT INTO students (id,firstname,lastname,age) VALUES ($1,$2,$3,$4)", student.Id, student.Firstname, student.Lastname, student.Age)
	return err
}
func (r *StudentRepository) Update(student *entity.Student) error {
	_, err := r.db.Exec("UPDATE students SET name = $2, age = $3 WHERE id = $1", student.Id, student.Name, student.Age)
	return err
}
func (r *StudentRepository) Delete(student *entity.Student) error {
	_, err := r.db.Exec("DELETE students  WHERE id = $1", student.Id)
	return err
}
