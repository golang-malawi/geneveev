package format

import (
	"fmt"
	"go/types"
)

type StructFormatter struct {
	TypeFormatterBase
}

func (f *StructFormatter) under(expr types.Type) *types.Struct {
	if namedType, ok := expr.(*types.Named); ok {
		if structType, ok := namedType.Underlying().(*types.Struct); ok {
			return structType
		}
	}
	return nil
}

func (f *StructFormatter) CanFormat(expr types.Type) bool {
	return f.under(expr) != nil
}

func (f *StructFormatter) Signature(expr types.Type) string {
	return expr.(*types.Named).Obj().Name()
}

func (f *StructFormatter) DefaultValue(_ types.Type) string {
	return ""
}

func (f *StructFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *StructFormatter) Constructor(fieldName string, _ types.Type) string {
	return "required this." + fieldName
}
