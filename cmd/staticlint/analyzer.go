package main

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

// MainExitAnalyzer search for inappropriate program exit
var MainExitAnalyzer = &analysis.Analyzer{
	Name: "mainExitCheck",
	Doc:  "check for unchecked errors",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.Name == "main" {
			inspectMainFunc(file, pass)
		}
	}
	return nil, nil
}

func inspectMainFunc(node ast.Node, pass *analysis.Pass) {
	ast.Inspect(node, func(node ast.Node) bool {
		if fn, ok := node.(*ast.FuncDecl); ok {

			if fn.Name.Name == "main" {
				inspectExitCall(fn, pass)
			}
		}
		return true
	})
}

func inspectExitCall(node ast.Node, pass *analysis.Pass) {
	ast.Inspect(node, func(node ast.Node) bool {
		if c, ok := node.(*ast.CallExpr); ok {
			if s, ok := c.Fun.(*ast.SelectorExpr); ok {
				if s.Sel.Name == "Exit" {
					pass.Reportf(s.Pos(), "os.Exit method call within main function")
				}
			}
		}
		return true
	})
}
