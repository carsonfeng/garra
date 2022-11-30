// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hello defines an Analyzer that checks if there is a hello param in func
package hello

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `checks if there is a hello param in func`

var Analyzer = &analysis.Analyzer{
	Name:     "hello",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var name, paramorder *bool

func init() {
	name = Analyzer.Flags.Bool("name", true, "will ensure context as parameter is named ctx")
	paramorder = Analyzer.Flags.Bool("paramorder", true, "will ensure context is the first argument of the functions")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		function := n.(*ast.FuncType)

		numContext := 0
		for _, f := range function.Params.List {
			t := pass.TypesInfo.TypeOf(f.Type)
			if `context.Context` == t.String() {
				numContext++
			}
		}
		if numContext > 1 {
			pass.Reportf(function.Pos(), "the function has more than one context defined in the parameters.")
			return
		}

		for key, f := range function.Params.List {
			// if the type is not a context, we do not need to go further

			// if it is a context we will check two things:
			// - it is the first parameter
			// - the variable is named ctx
			if *paramorder {
				if 0 != key {
					pass.Reportf(function.Pos(), "the function has a context that is not the first argument.")
				}
			}
			if *name {
				if len(f.Names) > 1 {
					continue
				}
				if `hello` != f.Names[0].Name {
					pass.Reportf(function.Pos(), "the function's first parameter that is not named 'hello'.")
				}
			}
		}
	})
	return nil, nil
}
