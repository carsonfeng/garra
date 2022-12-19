// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package package xorm_index_type_mismatch defines an Analyzer that checks xorm index type mismatch
package xorm_index_type_mismatch

import (
	"fmt"
	"github.com/carsonfeng/garra/common"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"strings"
)

const DEBUG = true
const Doc = `[Ziipin-Best-Practices] 字符串类型的字段索引，如果传入参数是int类型，索引不会生效。切记切记！[Sawa]`

var Analyzer = &analysis.Analyzer{
	Name:     "xorm_index_type_mismatch",
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

func checkCallObj(pass *analysis.Pass, call *ast.CallExpr, indexType map[string]string) {
	//check
	if selExpr, ok := call.Fun.(*ast.SelectorExpr); ok {
		funcName := selExpr.Sel.Name
		if "In" == funcName || "NotIn" == funcName {
			if len(call.Args) > 1 {
				if basicLit, ok2 := call.Args[0].(*ast.BasicLit); ok2 {
					field := strings.Trim(basicLit.Value, "\"")
					checkType := indexType[field]
					if checkType != "" {
						tav := pass.TypesInfo.Types[call.Args[1]]
						if sli, ok3 := tav.Type.(*types.Slice); ok3 {
							argType := sli.Elem().String()
							if argType != checkType {
								common.Reportf(pass, "Ziipin-Garra-XORM-Index-TypeMismatch", call.Args[1].Pos(),
									fmt.Sprintf("%s类型的字段索引(%s)，如果%s函数传入参数是%s类型，索引不会生效。", checkType, field, funcName, argType))
							}
						}
					}
				}
			}

		}
		if _callExp, ok2 := selExpr.X.(*ast.CallExpr); ok2 {
			checkCallObj(pass, _callExp, indexType)
		}
	}

}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		//(*ast.SelectorExpr)(nil),
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		callExpr := n.(*ast.CallExpr)
		if selectExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
			if !(selectExpr.Sel.Name == "Find" || selectExpr.Sel.Name == "FindAndCount" || selectExpr.Sel.Name == "Get") {
				return
			}
		} else {
			return
		}

		if len(callExpr.Args) == 0 {
			return
		}

		parseIndexType := func(callExpr *ast.CallExpr) (r map[string]string) {
			r = map[string]string{}
			tav := pass.TypesInfo.Types[callExpr.Args[0]]
			if pointer, ok := tav.Type.(*types.Pointer); !ok {
				return
			} else {
				switch x := pointer.Elem().(type) {
				case *types.Slice:
					if pointer2, ok2 := x.Elem().(*types.Pointer); ok2 {
						if named, ok3 := pointer2.Elem().(*types.Named); ok3 {
							if struct2, ok4 := named.Underlying().(*types.Struct); ok4 {
								for i := 0; i < struct2.NumFields(); i++ {
									lowerTag := strings.ToLower(struct2.Tag(i))
									if strings.Contains(lowerTag, "xorm") && strings.Contains(lowerTag, "index") {
										field := snakeString(struct2.Field(i).Name())
										typ := ""
										if strings.Contains(lowerTag, "varchar") {
											typ = "string"
										} else if strings.Contains(lowerTag, "int") {
											typ = "int"
										}
										if "" != typ {
											r[field] = typ
										}
									}
								}
							}

						}
					}
				}
			}
			return
		}

		indexType := parseIndexType(callExpr)
		if 0 == len(indexType) {
			return
		}

		checkCallObj(pass, callExpr, indexType)

	})

	return nil, nil
}
