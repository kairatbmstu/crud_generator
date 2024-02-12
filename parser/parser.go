package modelparser

import (
	"fmt"

	"example.com/ast1/model"
)

func ParseFile(file string) model.Model {
	return model.Model{}
}

func main() {
	entityDefs := []string{
		"entity Student {",
		"    id UUID",
		"    firstname String",
		"    lastname String",
		"    age String",
		"}",
		"",
		"entity Group {",
		"    id UUID",
		"    code String",
		"    startYear Integer",
		"}",
	}

	relationshipDefs := []string{
		"relationship ManyToOne {",
		"    Student{group} to Group",
		"}",
	}

	model := ParseModel(entityDefs, relationshipDefs)
	fmt.Printf("%+v\n", model)
}

func ParseModel(entityDefs, relationshipDefs []string) model.Model {
	var entities []model.Entity
	var relationships []model.Relationship

	for _, def := range entityDefs {
		if def == "" || def == "}" {
			continue
		}
		entityName := def[len("entity ") : len(def)-2] // Extracting entity name
		fields := make([]model.Field, 0)

		// Parsing fields for the entity
		for _, fieldDef := range entityDefs {
			if fieldDef == "" || fieldDef == "}" {
				continue
			}
			if fieldDef[0] == ' ' {
				fieldParts := parseField(fieldDef)
				field := model.Field{Name: fieldParts[0], Type: model.FieldType(fieldParts[1])}
				fields = append(fields, field)
			}
		}

		entity := model.Entity{Name: entityName, Fields: fields}
		entities = append(entities, entity)
	}

	for _, def := range relationshipDefs {
		if def == "" || def == "}" {
			continue
		}
		parts := parseRelationship(def)

		relationship := model.Relationship{
			EntityOne:      parts[0],
			EntityOneField: parts[1],
			EntityTwo:      parts[2],
			EntityTwoField: parts[3],
			RelationType:   model.RelationshipType(parts[4]),
		}
		relationships = append(relationships, relationship)
	}

	return model.Model{Entities: entities, Relationships: relationships}
}

func parseField(fieldDef string) []string {
	// Assuming the format is: "<name> <type>"
	return []string{} // Placeholder
}

func parseRelationship(relationshipDef string) []string {
	// Assuming the format is: "<EntityOne>{<EntityOneField>} to <EntityTwo>"
	return []string{} // Placeholder
}
