package mapper

import (
	"example.com/ast1/test/model"
	"example.com/ast1/test/repository"
)

type StudentMapper struct {
}

func (r *StudentMapper) ToDTO(student *model.Student) dto.StudentDTO {
	return err
}
func (r *StudentMapper) ToEntity(student *dto.StudentDTO) entity.Student {
	return err
}
