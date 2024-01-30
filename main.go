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
	outputFile string
)

func init() {
	flag.StringVar(&packageDir, "d", "", "directory to parse structs from")
	flag.StringVar(&outputDir, "output-dir", "", "directory to place generated code")
}

func main() {
	flag.Parse()

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, packageDir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, file := range packages {
		ast.Inspect(file, generateYupSchemas)
	}
}
