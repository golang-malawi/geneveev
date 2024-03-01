package format

import (
	"fmt"
	"go/types"
)

type ArrayFormatter struct {
	TypeFormatterBase
}

func (f *ArrayFormatter) under(expr types.Type) (TypeFormatter, types.Type) {
	sliceExpr := expr.(*types.Slice)
	elem := sliceExpr.Elem()
	formatter := f.Registry.GetTypeFormatter(elem)
	return formatter, elem
}

func (f *ArrayFormatter) CanFormat(expr types.Type) bool {
	_, ok := expr.(*types.Slice)
	return ok
}

func (f *ArrayFormatter) Signature(expr types.Type) string {
	formatter, expr := f.under(expr)
	return fmt.Sprintf("List<%s>", formatter.Signature(expr))
}

func (f *ArrayFormatter) DefaultValue(expr types.Type) string {
	return fmt.Sprintf("<%s>[]", f.Signature(expr))
}

func (f *ArrayFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *ArrayFormatter) Constructor(fieldName string, _ types.Type) string {
	return "required this." + fieldName
}
