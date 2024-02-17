package mapper

import (
	"example.com/ast1/test/model"
	"example.com/ast1/test/repository"
)

type GroupMapper struct {
}

func (r *GroupMapper) ToDTO(group *model.Group) dto.GroupDTO {
	return err
}
func (r *GroupMapper) ToEntity(group *dto.GroupDTO) entity.Group {
	return err
}
