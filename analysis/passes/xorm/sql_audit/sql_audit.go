// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package package sql_audit defines an Analyzer that audits xorm sql
package sql_audit

import (
	"fmt"
	"github.com/carsonfeng/garra/common"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"regexp"
	"strings"
)

const DEBUG = true
const Doc = `[Ziipin-Best-Practices] SQL审核：表update/delete操作没有指定要具体列，会导致全表变更，请核对`

var Analyzer = &analysis.Analyzer{
	Name:     "xorm_sql_audit",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	//Requires: []*analysis.Analyzer{buildssa.Analyzer},
	Run: run,
}

/**
 * snake string
 * @description XxYy to xx_yy , XxYY to xx_y_y
 * @param s input string
 * @return string
 **/
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func parseWhereArgType(pass *analysis.Pass, call *ast.CallExpr) (r map[string]string) {
	if nil == call || 0 == len(call.Args) {
		return
	}
	r = map[string]string{}

	getArgType := func(expr ast.Expr) (_r string) {
		switch x2 := expr.(type) {
		case *ast.BasicLit:
			if x2.Kind == token.INT {
				_r = "int"
			} else if x2.Kind == token.STRING {
				_r = "string"
			}
		case *ast.Ident:
			_r = pass.TypesInfo.Types[expr].Type.String()
		}
		return
	}

	switch x := call.Args[0].(type) {
	case *ast.CompositeLit:
		//eg. Where(builder.Eq{"room_id": 235, "uid": 33})
		for _, elt := range x.Elts {
			if kv, ok := elt.(*ast.KeyValueExpr); ok {
				key := ""

				if k, ok1 := kv.Key.(*ast.BasicLit); ok1 {
					key = strings.ReplaceAll(k.Value, "\"", "")
				} else {
					continue
				}
				valueType := getArgType(kv.Value)

				if valueType != "" {
					r[key] = valueType
				}
			}

		}
	case *ast.BasicLit:
		//eg. Where("room_id = ?", 123)
		if 2 != len(call.Args) {
			break
		}
		arg0 := x.Value

		inputVars := func(a string) (_r []string) {
			sli := regexp.MustCompile(`["=<>!\s+]`).Split(a, -1)
			var sli2 []string
			for _, item := range sli {
				if "" != item {
					sli2 = append(sli2, item)
				}
			}
			for i, item := range sli2 {
				if "?" == item && i > 0 {
					if v := sli2[i-1]; v != "" {
						_r = append(_r, v)
					}
				}
			}
			return
		}(arg0)

		if len(inputVars) > 0 && len(inputVars) == len(call.Args)-1 {
			for i := 1; i < len(call.Args); i++ {
				key := inputVars[i-1]
				if typ := getArgType(call.Args[i]); "" != typ {
					r[key] = typ
				}
			}
		}
	}
	return
}

var expFunc = map[string]bool{
	"In":    true,
	"NotIn": true,
	"Where": true,
	"And":   true,
	"Or":    true,
}

func checkCallObj(pass *analysis.Pass, call *ast.CallExpr) (hasCol bool) {
	//check
	if selExpr, ok := call.Fun.(*ast.SelectorExpr); ok {
		funcName := selExpr.Sel.Name

		if true == expFunc[funcName] {
			return true
		}

		if _callExp, ok2 := selExpr.X.(*ast.CallExpr); ok2 {
			hasCol = checkCallObj(pass, _callExp)
		}
	}
	return
}

var ops = map[string]int{"Update": 1, "Delete": 1}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		//(*ast.SelectorExpr)(nil),
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		callExpr := n.(*ast.CallExpr)

		if selectExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
			if ops[selectExpr.Sel.Name] == 0 {
				return
			}
		} else {
			return
		}

		if len(callExpr.Args) == 0 {
			return
		}

		isXormFunc := func(callExpr *ast.CallExpr) (r bool) {
			tav := pass.TypesInfo.Types[callExpr.Args[0]]
			if pointer, ok := tav.Type.(*types.Pointer); !ok {
				return false
			} else {
				procStruct := func(struct2 *types.Struct) {
					for i := 0; i < struct2.NumFields(); i++ {
						lowerTag := strings.ToLower(struct2.Tag(i))
						if strings.Contains(lowerTag, "xorm") {
							r = true
							return
						}
					}
				}

				switch x := pointer.Elem().(type) {

				case *types.Slice:
					if pointer2, ok2 := x.Elem().(*types.Pointer); ok2 {
						if named, ok3 := pointer2.Elem().(*types.Named); ok3 {
							if struct2, ok4 := named.Underlying().(*types.Struct); ok4 {
								procStruct(struct2)
							}
						}
					}

				case *types.Pointer:
					if named, ok2 := x.Elem().(*types.Named); ok2 {
						if struct2, ok3 := named.Underlying().(*types.Struct); ok3 {
							procStruct(struct2)
						}
					}
				}
			}
			return
		}

		if r := isXormFunc(callExpr); !r {
			return
		}

		if !checkCallObj(pass, callExpr) {
			common.Reportf(pass, "Ziipin-Garra-XORM-Sql-Audit", callExpr.Pos(),
				fmt.Sprintf(" SQL审核：表update/delete操作没有指定要具体列，会导致全表变更，请核对"))
		}

	})

	return nil, nil
}
