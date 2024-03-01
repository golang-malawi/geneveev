package format

import (
	"fmt"
	"go/types"
)

type AliasFormatter struct {
	TypeFormatterBase
}

func (f *AliasFormatter) under(expr types.Type) types.Type {
	if namedType, ok := expr.(*types.Named); ok {
		return namedType.Underlying()
	}

	return expr
}

func (f *AliasFormatter) CanFormat(expr types.Type) bool {
	return f.under(expr) != nil
}

func (f *AliasFormatter) Signature(expr types.Type) string {
	u := f.under(expr)
	return f.Registry.GetTypeFormatter(u).Signature(u)
}

func (f *AliasFormatter) DefaultValue(_ types.Type) string {
	return ""
}

func (f *AliasFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *AliasFormatter) Constructor(fieldName string, expr types.Type) string {
	u := f.under(expr)
	return f.Registry.GetTypeFormatter(u).Constructor(fieldName, u)
}
