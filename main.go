package main

import (
	"fmt"
	"io/ioutil"

	"example.com/ast1/codegenerator"
	"example.com/ast1/parser"
)

func main() {
	// // Create a new file set.

	data, err := ioutil.ReadFile("parser/example.jdl")
	if err != nil {
		panic(err)
	}
	jdlText := string(data)

	tokens, err := parser.Tokenize(jdlText)

	if err != nil {
		panic(err)
	}

	err = parser.LexicalAnalysis(tokens)
	if err != nil {
		panic(err)
	}

	model, err := parser.ParseModel(tokens)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)

	for _, ent := range model.Entities {
		codegenerator.GenerateEntity("test/entity", &ent)
		codegenerator.GenerateDTO("test/dto", &ent)
		codegenerator.GenerateMapper("test/mapper", &ent)
		codegenerator.GenerateRepository("test/repository", &ent)
		codegenerator.GenerateService("test/service", &ent)
		codegenerator.GenerateRestApiHandler("test/handler", &ent)
	}

}
