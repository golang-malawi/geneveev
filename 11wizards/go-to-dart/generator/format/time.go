package format

import (
	"fmt"
	"go/types"
)

type TimeFormatter struct {
	TypeFormatterBase
}

func (f *TimeFormatter) CanFormat(expr types.Type) bool {
	if namedType, ok := expr.(*types.Named); ok {
		if namedType.Obj().Type().String() == "time.Time" {
			return true
		}
	}

	return false
}

func (f *TimeFormatter) Signature(_ types.Type) string {
	return "DateTime"
}

func (f *TimeFormatter) DefaultValue(_ types.Type) string {
	return ""
}

func (f *TimeFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *TimeFormatter) Constructor(fieldName string, _ types.Type) string {
	return fmt.Sprintf("required this.%s", fieldName)
}
