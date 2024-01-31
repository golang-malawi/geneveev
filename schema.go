package main

import (
	"fmt"
	"strings"
)

type Schema struct {
	name   string
	Fields map[string]string
}

func NewSchema(name string) *Schema {
	return &Schema{
		name:   name,
		Fields: make(map[string]string),
	}
}

func (s *Schema) addEntry(name, expr string) {
	if name != "" && expr != "" {
		s.Fields[name] = expr
	}
}

func (s *Schema) AddBoolField(name, yupExpression string) {
	s.addEntry(name, yupExpression)
}

func (s *Schema) AddStringField(name, yupExpression string) {
	s.addEntry(name, yupExpression)
}

func (s *Schema) AddNumberField(name, yupExpression string) {
	s.addEntry(name, yupExpression)
}

func (s *Schema) AddTimeField(name, yupExpression string) {
	s.addEntry(name, yupExpression)
}

func (s *Schema) AddMixedField(name, yupExpression string) {
	s.addEntry(name, yupExpression)
}

func (s *Schema) Name() string {
	return s.name + "Schema"
}

func (s *Schema) Filename(snakecase bool) string {
	if snakecase {
		return strings.ToLower(s.name + "_schema" + ".js")
	}

	return strings.ToUpper(string(s.Name()[0])) + s.Name()[1:] + ".js"
}

func (s *Schema) IsEmpty() bool {
	return len(s.Fields) < 1
}

func (s *Schema) ToJS() string {
	var sb strings.Builder
	sb.WriteString("export const ")
	sb.WriteString(s.Name())
	sb.WriteString(" = yup.object({\n")

	for fieldName, expr := range s.Fields {
		sb.WriteString(fmt.Sprintf("\t%s: %s,\n", fieldName, expr))
	}

	sb.WriteString("})\n")

	return sb.String()
}
