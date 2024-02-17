package repository

import (
	"database/sql"
	"example.com/ast1/test/model"
	"example.com/ast1/test/repository"
)

type GroupRepository struct {
	db *sql.DB
}

func (r *GroupRepository) Create(group *model.Group) error {
	_, err := r.db.Exec("INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}
func (r *GroupRepository) Update(group *model.Group) error {
	_, err := r.db.Exec("UPDATE students SET name = $2, age = $3 WHERE id = $1", student.Id, student.Name, student.Age)
	return err
}
func (r *GroupRepository) Delete(group *model.Group) error {
	_, err := r.db.Exec("DELETE students  WHERE id = $1", student.Id)
	return err
}
