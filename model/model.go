package model

const (
	FieldType_int     = "int"
	FieldType_string  = "string"
	FieldType_uuid    = "uuid"
	FieldType_boolean = "bool"
)

type FieldType string

type Field struct {
	Name       string
	Type       FieldType
	IsRequired bool
}

type Model struct {
	Name   string
	Fields []Field
}

type RelationType string

const (
	RelationType_OneToOne   = "onetoone"
	RelationType_OneToMany  = "onetomany"
	RelationType_ManyToOne  = "manytoone"
	RelationType_ManyToMany = "manytomany"
)

type Relation struct {
	ModelOne     Model
	ModelTwo     Model
	RelationType RelationType
}

type File struct {
	Name string
}
