package mapper

import (
	"example.com/ast1/test/entity"
	"example.com/ast1/test/dto"
)

type StudentMapper struct {
}

func (r *StudentMapper) ToDTO(student *entity.Student) dto.StudentDTO {
	var studentDTO = dto.StudentDTO{}
	studentDTO.Id = student.Id
	studentDTO.Firstname = student.Firstname
	studentDTO.Lastname = student.Lastname
	studentDTO.Age = student.Age
	return studentDTO
}
func (r *StudentMapper) ToEntity(studentDTO *dto.StudentDTO) entity.Student {
	var student = entity.Student{}
	return student
}
