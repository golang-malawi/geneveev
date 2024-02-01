package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
)

var (
	packageDir string
	outputDir  string
	snakeCase  bool
	zodMode    bool
)

func init() {
	flag.StringVar(&packageDir, "d", "", "directory to parse structs from")
	flag.StringVar(&outputDir, "output-dir", "", "directory to place generated code")
	flag.BoolVar(&snakeCase, "snakecase", false, "whether to name files as snake_case")
	flag.BoolVar(&zodMode, "zod", false, "whether to generate zod schemas or not")
}

func main() {
	flag.Parse()

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, packageDir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, file := range packages {
		if zodMode {
			ast.Inspect(file, generator(Zod()))
		} else {
			ast.Inspect(file, generator(Yup()))
		}
	}
}
