package mo

import (
	"fmt"
	"go/types"

	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/format"
)

type OptionFormatter struct {
	format.TypeFormatterBase
}

func (f *OptionFormatter) under(expr types.Type) (format.TypeFormatter, types.Type) {
	e := expr.(*types.Named).TypeArgs().At(0)
	formatter := f.Registry.GetTypeFormatter(e)
	return formatter, e
}

func (f *OptionFormatter) CanFormat(expr types.Type) bool {
	if namedType, ok := expr.(*types.Named); ok {
		if namedType.Obj().Type().String() == "github.com/samber/mo.Option[T any]" {
			return true
		}
	}

	return false
}

func (f *OptionFormatter) Signature(expr types.Type) string {
	formatter, expr := f.under(expr)
	return fmt.Sprintf("%s?", formatter.Signature(expr))
}

func (f *OptionFormatter) DefaultValue(_ types.Type) string {
	return ""
}

func (f *OptionFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *OptionFormatter) Constructor(fieldName string, _ types.Type) string {
	return "this." + fieldName
}
