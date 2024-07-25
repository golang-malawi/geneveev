package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/alecthomas/kong"
	dartGenerator "github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator"
	"github.com/golang-malawi/geneveev/11wizards/go-to-dart/generator/options"
)

type Context struct {
	Debug bool
}

type LsCmd struct {
	Paths []string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
}

type GenerateCmd struct {
	Format     string `arg:"" type:"string" help:"Which format to generate to"`
	PackageDir string `name:"d" type:"path" help:"package containing structs to generate from"`
	OutputDir  string `name:"" type:"path"`
	Snakecase  bool   `name:"" help:"whether to name files as snake_case"`
}

func (g *GenerateCmd) Run(ctx *Context) error {
	packageDir := g.PackageDir
	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, packageDir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, file := range packages {
		switch g.Format {
		case "dart":
			o := options.Options{
				Input:   packageDir,
				Output:  g.OutputDir,
				Imports: []string{},
				Mode:    options.Mode(options.JSON),
			}

			if o.Mode != options.JSON && o.Mode != options.Firestore {
				fmt.Println("Mode must be either json or firestore")
				os.Exit(1)
			}
			dartGenerator.Run(o)
			break
		case "zod":
			ast.Inspect(file, generator(Zod(), g.Snakecase, g.OutputDir))
			break
		case "yup":
			ast.Inspect(file, generator(Yup(), g.Snakecase, g.OutputDir))
			break
		}
	}

	return nil
}

var cli struct {
	Debug    bool        `help:"Enable debug mode."`
	Generate GenerateCmd `cmd:"" help:"Generate subcommand to generate different formats."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
