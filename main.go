package main

import (
	"example.com/ast1/codegenerator"
	"example.com/ast1/model"
)

func main() {
	// // Create a new file set.

	codegenerator.GenerateEntity("test/model", &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			model.Field{
				Name: "id",
				Type: model.FieldType_uuid,
			},
			model.Field{
				Name: "name",
				Type: model.FieldType_string,
			},
			model.Field{
				Name: "age",
				Type: model.FieldType_int,
			},
		},
	})

	codegenerator.GenerateDTO("test/dto", &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			model.Field{
				Name: "id",
				Type: model.FieldType_uuid,
			},
			model.Field{
				Name: "name",
				Type: model.FieldType_string,
			},
			model.Field{
				Name: "age",
				Type: model.FieldType_int,
			},
		},
	})

	codegenerator.GenerateMapper("test/mapper", &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			model.Field{
				Name: "id",
				Type: model.FieldType_uuid,
			},
			model.Field{
				Name: "name",
				Type: model.FieldType_string,
			},
			model.Field{
				Name: "age",
				Type: model.FieldType_int,
			},
		},
	})

	codegenerator.GenerateRepository("test/repository", &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			model.Field{
				Name: "id",
				Type: model.FieldType_uuid,
			},
			model.Field{
				Name: "name",
				Type: model.FieldType_string,
			},
			model.Field{
				Name: "age",
				Type: model.FieldType_int,
			},
		},
	})

	codegenerator.GenerateService("test/service", &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			model.Field{
				Name: "id",
				Type: model.FieldType_uuid,
			},
			model.Field{
				Name: "name",
				Type: model.FieldType_string,
			},
			model.Field{
				Name: "age",
				Type: model.FieldType_int,
			},
		},
	})

	codegenerator.GenerateRestApiHandler("test/handler", &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			model.Field{
				Name: "id",
				Type: model.FieldType_uuid,
			},
			model.Field{
				Name: "name",
				Type: model.FieldType_string,
			},
			model.Field{
				Name: "age",
				Type: model.FieldType_int,
			},
		},
	})
}
