package format

import (
	"fmt"
	"go/types"
)

type PointerFormatter struct {
	TypeFormatterBase
}

func (f *PointerFormatter) under(expr types.Type) (TypeFormatter, types.Type) {
	starExpr := expr.(*types.Pointer)
	formatter := f.Registry.GetTypeFormatter(starExpr.Elem())
	return formatter, starExpr.Elem()
}

func (f *PointerFormatter) CanFormat(expr types.Type) bool {
	_, ok := expr.(*types.Pointer)
	return ok
}

func (f *PointerFormatter) Signature(expr types.Type) string {
	formatter, expr := f.under(expr)
	return fmt.Sprintf("%s?", formatter.Signature(expr))
}

func (f *PointerFormatter) DefaultValue(_ types.Type) string {
	return ""
}

func (f *PointerFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *PointerFormatter) Constructor(fieldName string, _ types.Type) string {
	return "this." + fieldName
}
