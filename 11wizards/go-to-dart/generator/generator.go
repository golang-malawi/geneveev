package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/format"
	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/format/mo"
	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/options"

	"golang.org/x/tools/go/packages"
)

//go:embed dart/timestamp_converter.dart
var timestampConverterSrc string

func generateHeader(pkg *packages.Package, wr io.Writer, mode options.Mode, imports []string) {
	if mode == options.Firestore {
		fmt.Fprint(wr, "import 'package:cloud_firestore/cloud_firestore.dart';\n")
	}

	fmt.Fprint(wr, "import 'package:json_annotation/json_annotation.dart';\n")
	for _, imp := range imports {
		fmt.Fprintf(wr, "import '%s';\n", imp)
	}

	fmt.Fprintf(wr, "\npart '%s.go.g.dart';\n\n", pkg.Name)

	if mode == options.Firestore {
		fmt.Fprint(wr, timestampConverterSrc)
		fmt.Fprint(wr, "\n\n")
	}
}

func createRegistry(mode options.Mode) *format.TypeFormatterRegistry {
	registry := format.NewTypeFormatterRegistry()

	base := format.TypeFormatterBase{Mode: mode}

	registry.RegisterTypeFormatter(&format.AliasFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.StructFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.PrimitiveFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.TimeFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.PointerFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.ArrayFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.MapFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&mo.OptionFormatter{TypeFormatterBase: base})

	return registry
}

func generateClasses(pkg *packages.Package, wr io.Writer, mode options.Mode) {
	registry := createRegistry(mode)

	for _, value := range pkg.TypesInfo.Defs {
		if value == nil {
			continue
		}

		if typeName, ok := value.(*types.TypeName); ok {
			registry.KnownTypes[typeName.Type()] = struct{}{}
		}
	}

	var list []struct {
		TypeName   *types.TypeName
		StructType *types.Struct
	}

	for _, value := range pkg.TypesInfo.Defs {
		if value == nil {
			continue
		}

		if typeName, ok := value.(*types.TypeName); ok {
			if structType, ok := typeName.Type().Underlying().(*types.Struct); ok {
				list = append(list, struct {
					TypeName   *types.TypeName
					StructType *types.Struct
				}{
					TypeName:   typeName,
					StructType: structType,
				})
			}
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].TypeName.Name() < list[j].TypeName.Name()
	})

	for _, item := range list {
		generateDartClass(wr, item.TypeName, item.StructType, registry, mode)
	}
}

func writeOut(output, outputDartFile string, wr *bytes.Buffer) {
	if _, err := os.Stat(output); os.IsNotExist(err) {
		err = os.MkdirAll(output, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	outputFilePath := filepath.Join(output, outputDartFile)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}

	defer func() { _ = outputFile.Close() }()

	_, err = outputFile.Write(wr.Bytes())

	if err != nil {
		panic(err)
	}

	fmt.Printf("Processed: %s -> %s\n", outputDartFile, outputFilePath)
}

func Run(options options.Options) {
	if abs, err := filepath.Abs(options.Input); err == nil {
		options.Input = abs
	} else {
		panic(err)
	}

	pkgs, err := packages.Load(&packages.Config{
		Dir:  options.Input,
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}, options.Input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			for _, err := range pkg.Errors {
				fmt.Println(err)
			}

			os.Exit(1)
		}
	}

	for _, pkg := range pkgs {
		var buf []byte
		wr := bytes.NewBuffer(buf)
		generateHeader(pkg, wr, options.Mode, options.Imports)
		generateClasses(pkg, wr, options.Mode)
		writeOut(options.Output, fmt.Sprintf("%s.go.dart", pkg.Name), wr)
	}
}
