package generator

import (
	"fmt"
	"go/types"
	"io"

	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/format"
	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/options"
	"github.com/openconfig/goyang/pkg/indent"
)

func generateFields(wr io.Writer, st *types.Struct, registry *format.TypeFormatterRegistry, mode options.Mode) {
	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		tag := st.Tag(i)
		generateFieldDeclaration(wr, field, tag, registry, mode)
		fmt.Fprintln(wr, ";")
	}
	fmt.Fprintln(wr)
}

func generateConstructor(wr io.Writer, ts *types.TypeName, st *types.Struct, registry *format.TypeFormatterRegistry) {
	fmt.Fprintf(wr, "%s({\n", ts.Name())

	for i := 0; i < st.NumFields(); i++ {
		f := st.Field(i)
		generateFieldConstrutor(indent.NewWriter(wr, "\t"), f, registry)
		fmt.Fprintln(wr, ",")
	}

	fmt.Fprintf(wr, "});")
	fmt.Fprintln(wr)
	fmt.Fprintln(wr)
}

func generateSerialization(wr io.Writer, ts *types.TypeName) {
	fmt.Fprintf(wr, "Map<String, dynamic> toJson() => _$%sToJson(this);\n\n", ts.Name())
}

func generateDeserialization(wr io.Writer, ts *types.TypeName) {
	fmt.Fprintf(wr, "factory %s.fromJson(Map<String, dynamic> json) => _$%sFromJson(json);\n", ts.Name(), ts.Name())
}

func generateDartClass(outputFile io.Writer, ts *types.TypeName, st *types.Struct, registry *format.TypeFormatterRegistry, mode options.Mode) {
	fmt.Fprintln(outputFile, "@JsonSerializable(explicitToJson: true)")
	if mode == options.Firestore {
		fmt.Fprintln(outputFile, "@_TimestampConverter()")
	}

	fmt.Fprintf(outputFile, "class %s {\n", ts.Name())

	wr := indent.NewWriter(outputFile, "\t")

	generateFields(wr, st, registry, mode)
	generateConstructor(wr, ts, st, registry)
	generateSerialization(wr, ts)
	generateDeserialization(wr, ts)

	fmt.Fprintln(outputFile, "}")
	fmt.Fprintln(outputFile, "")
}
