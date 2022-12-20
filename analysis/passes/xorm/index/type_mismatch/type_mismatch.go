// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package package xorm_index_type_mismatch defines an Analyzer that checks xorm index type mismatch
package xorm_index_type_mismatch

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
							if checkType == "string" && argType == "int" {
								common.Reportf(pass, "Ziipin-Garra-XORM-Index-TypeMismatch", call.Args[1].Pos(),
									fmt.Sprintf("%s类型的字段索引(%s)，%s函数传入参数是%s类型，索引不会生效。", checkType, field, funcName, argType))
							}
						}
					}
				}
			}
		} else if "Where" == funcName || "And" == funcName || "Or" == funcName {
			if argTyp := parseWhereArgType(pass, call); len(argTyp) > 0 {
				for k, v := range argTyp {
					if "string" == indexType[k] && "int" == v {
						common.Reportf(pass, "Ziipin-Garra-XORM-Index-TypeMismatch", call.Pos(),
							fmt.Sprintf("%s类型的字段索引(%s)，%s函数传入参数是%s类型，索引不会生效。", indexType[k], k, funcName, v))
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
