package format

import (
	"fmt"
	"go/types"

	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/options"
)

type TypeFormatter interface {
	SetRegistry(registry *TypeFormatterRegistry)
	CanFormat(expr types.Type) bool
	Signature(expr types.Type) string
	DefaultValue(expr types.Type) string
	Declaration(fieldName string, expr types.Type) string
	Constructor(fieldName string, expr types.Type) string
}

type TypeFormatterBase struct {
	Registry *TypeFormatterRegistry
	Mode     options.Mode
}

func (t *TypeFormatterBase) SetRegistry(registry *TypeFormatterRegistry) {
	t.Registry = registry
}

type TypeFormatterRegistry struct {
	KnownTypes map[types.Type]struct{}
	Formatters []TypeFormatter
}

func NewTypeFormatterRegistry() *TypeFormatterRegistry {
	return &TypeFormatterRegistry{
		KnownTypes: make(map[types.Type]struct{}),
		Formatters: make([]TypeFormatter, 0),
	}
}
func (t *TypeFormatterRegistry) RegisterTypeFormatter(formatter TypeFormatter) {
	t.Formatters = append(t.Formatters, formatter)
	formatter.SetRegistry(t)
}

func (t *TypeFormatterRegistry) GetTypeFormatter(expr types.Type) TypeFormatter {
	// walks the t.Formatters in reverse order
	// so that the last registered formatter is the first to be checked
	for i := len(t.Formatters) - 1; i >= 0; i-- {
		f := t.Formatters[i]
		if f.CanFormat(expr) {
			return f
		}
	}

	panic(fmt.Sprintf("no formatter found for %v", expr))
}
