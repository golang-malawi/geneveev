package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	dartGenerator "github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator"
	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/options"
)

var (
	packageDir string
	outputDir  string
	snakeCase  bool
	zodMode    bool
	dartMode   bool
)

func init() {
	flag.StringVar(&packageDir, "d", "", "directory to parse structs from")
	flag.StringVar(&outputDir, "output-dir", "", "directory to place generated code")
	flag.BoolVar(&snakeCase, "snakecase", false, "whether to name files as snake_case")
	flag.BoolVar(&zodMode, "zod", false, "whether to generate zod schemas or not")
	flag.BoolVar(&dartMode, "dart", false, "whether to generate Dart classes or not")
}

func main() {
	flag.Parse()

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, packageDir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, file := range packages {
		if dartMode {

			o := options.Options{
				Input:   packageDir,
				Output:  outputDir,
				Imports: []string{},
				Mode:    options.Mode(options.JSON),
			}

			if o.Mode != options.JSON && o.Mode != options.Firestore {
				fmt.Println("Mode must be either json or firestore")
				os.Exit(1)
			}
			dartGenerator.Run(o)
		}

		if zodMode {
			ast.Inspect(file, generator(Zod()))
		} else {
			ast.Inspect(file, generator(Yup()))
		}
	}
}
