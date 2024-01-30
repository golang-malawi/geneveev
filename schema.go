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

func (s *Schema) AddBoolField(name, yupExpression string) {
	s.Fields[name] = yupExpression
}

func (s *Schema) AddStringField(name, yupExpression string) {
	s.Fields[name] = yupExpression
}

func (s *Schema) AddNumberField(name, yupExpression string) {
	s.Fields[name] = yupExpression
}

func (s *Schema) AddTimeField(name, yupExpression string) {
	s.Fields[name] = yupExpression
}

func (s *Schema) AddMixedField(name, yupExpression string) {
	s.Fields[name] = yupExpression
}

func (s *Schema) Name() string {
	return s.name + "Schema"
}

func (s *Schema) Filename() string {
	return strings.ToLower(s.name + "_schema" + ".js")
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
