package format

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"

	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/options"
	"github.com/iancoleman/strcase"
)

func GetFieldName(f *types.Var) string {
	if f.Anonymous() {
		panic(fmt.Sprintf("no name for field: %#v", f))
	}

	return strcase.ToLowerCamel(f.Name())
}

func GetJSONFieldName(tag string, mode options.Mode) string {
	var tagName string
	if mode == options.Firestore {
		tagName = "firestore"
	} else {
		tagName = "json"
	}

	if tag != "" {
		val := reflect.StructTag(strings.Trim(tag, "`"))
		value, ok := val.Lookup(tagName)
		if ok {
			return strings.Split(value, ",")[0]
		}
	}

	return ""
}
