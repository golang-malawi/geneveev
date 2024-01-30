package main

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

const (
	intTypeRegex   = `^[u]?int\d?`
	floatTypeRegex = `^float\d?`
)

func isIntegerType(field *ast.Field) bool {
	if field == nil {
		return false
	}

	match, err := regexp.MatchString(intTypeRegex, fmt.Sprint(field.Type))
	if err != nil {
		return false
	}
	return match
}

func isFloatType(field *ast.Field) bool {
	if field == nil {
		return false
	}

	match, err := regexp.MatchString(floatTypeRegex, fmt.Sprint(field.Type))
	if err != nil {
		return false
	}
	return match
}

func isTimeField(field *ast.Field) bool {
	if field == nil {
		return false
	}
	return fmt.Sprint(field.Type) == "&{time Time}"
}

func generateYupSchemas(n ast.Node) bool {
	t, ok := n.(*ast.TypeSpec)
	if !ok {
		return true
	}

	s, ok := t.Type.(*ast.StructType)
	if !ok {
		return true
	}

	var sb strings.Builder
	sb.WriteString("const ")
	sb.WriteString(t.Name.Name + "Schema")
	sb.WriteString(" = yup.object({\n")

	// fmt.Println("DEBUG: processing type", t.Name.Name)
	for _, field := range s.Fields.List {
		if field.Tag == nil {
			continue
		}

		start := strings.Index(field.Tag.Value, "validate:")
		if start <= 0 {
			continue
		}

		validateTag := field.Tag.Value[start:]
		end := strings.LastIndex(validateTag, `"`)
		if end != -1 {
			validateTag = validateTag[:end]
		}

		var mapped string
		if fmt.Sprint(field.Type) == "string" {
			mapped = mapStringTag(validateTag)
		} else if fmt.Sprint(field.Type) == "bool" {
			mapped = mapBoolTag(validateTag)
		} else if isIntegerType(field) || isFloatType(field) {
			mapped = mapNumberTag(validateTag)
		} else if isTimeField(field) {
			mapped = mapTimeStructTag(validateTag)
		} else {
			// default to using a "mixed" field
			mapped = mapMixedFieldTag(validateTag)
		}

		sb.WriteString(fmt.Sprintf("\t%s: %s,\n", field.Names[0].Name, mapped))
	}
	sb.WriteString("})\n")

	fmt.Println(sb.String())
	return false
}
