package format

import (
	"fmt"
	"go/types"
)

type MapFormatter struct {
	TypeFormatterBase
}

func (f *MapFormatter) under(expr types.Type) (TypeFormatter, TypeFormatter, types.Type, types.Type) {
	mapExpr := expr.(*types.Map)
	keyFormatter := f.Registry.GetTypeFormatter(mapExpr.Key())
	valueFormatter := f.Registry.GetTypeFormatter(mapExpr.Elem())
	return keyFormatter, valueFormatter, mapExpr.Key(), mapExpr.Elem()
}

func (f *MapFormatter) CanFormat(expr types.Type) bool {
	_, ok := expr.(*types.Map)
	return ok
}

func (f *MapFormatter) Signature(expr types.Type) string {
	keyFormatter, valueFormatter, keyExpr, valueExpr := f.under(expr)
	return fmt.Sprintf("Map<%s, %s>", keyFormatter.Signature(keyExpr), valueFormatter.Signature(valueExpr))
}

func (f *MapFormatter) DefaultValue(expr types.Type) string {
	keyFormatter, valueFormatter, keyExpr, valueExpr := f.under(expr)
	return fmt.Sprintf("<%s, %s>{}", keyFormatter.Signature(keyExpr), valueFormatter.Signature(valueExpr))
}

func (f *MapFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *MapFormatter) Constructor(fieldName string, _ types.Type) string {
	return "required this." + fieldName
}
