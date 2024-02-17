package repository

import (
	"database/sql"
	"example.com/ast1/test/entity"
)

type GroupRepository struct {
	db *sql.DB
}

func (r *GroupRepository) Create(group *entity.Group) error {
	_, err := r.db.Exec("INSERT INTO groups (id,code,startyear) VALUES ($1,$2,$3)", group.Id, group.Code, group.StartYear)
	return err
}
func (r *GroupRepository) Update(group *entity.Group) error {
	_, err := r.db.Exec("UPDATE  groups set code = $2,startyear = $3, WHERE id = $1", group.Id, group.Code, group.StartYear)
	return err
}
func (r *GroupRepository) Delete(group *entity.Group) error {
	_, err := r.db.Exec("DELETE groups  WHERE id = $1", group.Id)
	return err
}
