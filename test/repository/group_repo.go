package repository

import (
	"database/sql"

	"example.com/ast1/test/entity"
)

type GroupRepository struct {
	db *sql.DB
}

func (r *GroupRepository) Create(group *entity.Group) error {
	_, err := r.db.Exec("INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}
func (r *GroupRepository) Update(group *entity.Group) error {
	_, err := r.db.Exec("UPDATE students SET name = $2, age = $3 WHERE id = $1", student.Id, student.Name, student.Age)
	return err
}
func (r *GroupRepository) Delete(group *entity.Group) error {
	_, err := r.db.Exec("DELETE students  WHERE id = $1", student.Id)
	return err
}
