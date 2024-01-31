package main

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
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

	schema := NewSchema(t.Name.Name)

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

		fieldName := field.Names[0].Name
		if fmt.Sprint(field.Type) == "string" {
			schema.AddStringField(fieldName, mapStringTag(validateTag))
		} else if fmt.Sprint(field.Type) == "bool" {
			schema.AddBoolField(fieldName, mapBoolTag(validateTag))
		} else if isIntegerType(field) || isFloatType(field) {
			schema.AddNumberField(fieldName, mapNumberTag(validateTag))
		} else if isTimeField(field) {
			schema.AddTimeField(fieldName, mapTimeStructTag(validateTag))
		} else {
			// default to using a "mixed" field
			schema.AddMixedField(fieldName, mapMixedFieldTag(validateTag))
		}
	}

	if schema.IsEmpty() {
		return false
	}

	if outputDir != "" {
		err := os.Mkdir(filepath.Join(outputDir), 0o775)
		if err != nil {
			if !errors.Is(err, os.ErrExist) {
				panic(err)
			}
		}
		err = os.WriteFile(filepath.Join(outputDir, schema.Filename(snakeCase)), []byte(schema.ToJSFile()), os.FileMode(0o777))
		if err != nil {
			panic(err)
		}
		return false
	}

	fmt.Println(schema.ToJS())
	return false
}
