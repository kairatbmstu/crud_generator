package mapper

import (
	"example.com/ast1/test/entity"
	"example.com/ast1/test/dto"
)

type GroupMapper struct {
}

func (r *GroupMapper) ToDTO(group *entity.Group) dto.GroupDTO {
	var groupDTO = dto.GroupDTO{}
	groupDTO.Id = group.Id
	groupDTO.Code = group.Code
	groupDTO.StartYear = group.StartYear
	return groupDTO
}
func (r *GroupMapper) ToEntity(groupDTO *dto.GroupDTO) entity.Group {
	var group = entity.Group{}
	return group
}
