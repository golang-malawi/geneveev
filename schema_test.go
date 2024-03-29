package main

import (
	"testing"
)

func TestNewSchema(t *testing.T) {
	schema := NewSchema("test")
	if schema.name != "test" {
		t.Error("expected schema to be named 'test'")
	}
}

func TestAddEntry(t *testing.T) {
	schema := NewSchema("test")

	schema.addEntry("fieldName", "yup.bool()")

	if len(schema.Fields) < 1 {
		t.Error("failed to find any fields after call to 'addEntry'")
	}
}

func TestName(t *testing.T) {
	schema := NewSchema("test")

	if schema.Name() != "testSchema" {
		t.Error("expected Schema.Name() to return 'testSchema'")
	}
}

func TestFilename(t *testing.T) {
	schema := NewSchema("test")

	if schema.Filename(true) != "test_schema.js" {
		t.Error("expected Schema.Filename() to return 'test_schema.js'")
	}

	if schema.Filename(false) != "TestSchema.js" {
		t.Error("expected Schema.Filename() to return 'TestSchema.js'")
	}
}

func TestToJS(t *testing.T) {
	schema := NewSchema("test")

	schema.addEntry("fieldName", "yup.bool()")

	expected := "const testSchema = yup.object({\n\tfieldName: yup.bool(),\n})\n"

	if expected != schema.ToJS(Yup()) {
		t.Errorf("failed to match \nwant=%s\nhave=%s", expected, schema.ToJS(Yup()))
	}
}

func TestToJSFile(t *testing.T) {
	schema := NewSchema("test")

	schema.addEntry("fieldName", "yup.bool()")

	expected := "import * as yup from 'yup';\n\nconst testSchema = yup.object({\n\tfieldName: yup.bool(),\n})\n\nexport default testSchema;\n"

	if expected != schema.ToJSFile(Yup()) {
		t.Errorf("failed to match \nwant=%s\nhave=%s", expected, schema.ToJSFile(Yup()))
	}

}
