package main

import "testing"

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

	if schema.Filename(false) != "testSchema.js" {
		t.Error("expected Schema.Filename() to return 'testSchema.js'")
	}
}

func TestToJS(t *testing.T) {
	schema := NewSchema("test")

	schema.addEntry("fieldName", "yup.bool()")

	expected := "export const testSchema = yup.object({\n\tfieldName: yup.bool(),\n})\n"

	if expected != schema.ToJS() {
		t.Errorf("failed to match \nwant=%s\nhave=%s", expected, schema.ToJS())
	}

}
