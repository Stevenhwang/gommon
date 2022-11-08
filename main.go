package main

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	dirs, _ := os.ReadDir(pwd)
	fset := token.NewFileSet()
	for _, x := range dirs {
		if x.IsDir() {
			d, _ := parser.ParseDir(fset, fmt.Sprintf("./%s", x.Name()), nil, parser.ParseComments)
			for _, f := range d {
				p := doc.New(f, "./", 2)
				for _, f := range p.Funcs {
					fmt.Println(f.Doc)
				}
			}
		}
	}
}
