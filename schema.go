package main

import (
	"fmt"
	"strings"
)

type Schema struct {
	name   string
	Fields []string
}

func NewSchema(name string) *Schema {
	return &Schema{
		name:   name,
		Fields: make([]string, 0),
	}
}

func (s *Schema) addEntry(name, expr string) {
	if name != "" && expr != "" {
		s.Fields = append(s.Fields, fmt.Sprintf("%s: %s,", name, expr))
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

func (s *Schema) ToJS(m mapper) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("const %s = %s({\n", s.Name(), m.object()))

	for _, expr := range s.Fields {
		sb.WriteString(fmt.Sprintf("\t%s\n", expr))
	}

	sb.WriteString("})\n")

	return sb.String()
}

func (s *Schema) ToJSFile(m mapper) string {
	var sb strings.Builder
	sb.WriteString(m.jsImport() + ";\n\n")
	sb.WriteString(s.ToJS(m))
	sb.WriteString("\nexport default ")
	sb.WriteString(s.Name() + ";\n")

	return sb.String()
}
