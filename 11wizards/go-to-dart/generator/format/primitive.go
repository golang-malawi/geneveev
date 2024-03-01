package format

import (
	"fmt"
	"go/types"
)

type PrimitiveFormatter struct {
	TypeFormatterBase
}

func (f *PrimitiveFormatter) toDartPrimitive(expr types.Type) string {
	if e, ok := expr.(*types.Basic); ok {
		switch e.Kind() {
		case types.Bool:
			return "bool"
		case types.Float32:
			return "double"
		case types.Float64:
			return "double"
		case types.Int:
			return "int"
		case types.Int16:
			return "int"
		case types.Int32:
			return "int"
		case types.Int64:
			return "int"
		case types.Int8:
			return "int"
		case types.String:
			return "String"
		case types.Uint:
			return "int"
		case types.Uint16:
			return "int"
		case types.Uint32:
			return "int"
		case types.Uint64:
			return "int"
		case types.Uint8:
			return "int"
		case types.Uintptr:
			return "int"
		}
	}

	return ""
}

func (f *PrimitiveFormatter) CanFormat(expr types.Type) bool {
	return f.toDartPrimitive(expr) != ""
}

func (f *PrimitiveFormatter) Signature(expr types.Type) string {
	return f.toDartPrimitive(expr)
}

func (f *PrimitiveFormatter) DefaultValue(_ types.Type) string {
	return ""
}

func (f *PrimitiveFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *PrimitiveFormatter) Constructor(fieldName string, _ types.Type) string {
	return "required this." + fieldName
}
