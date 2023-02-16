// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package package sql_audit defines an Analyzer that audits xorm sql
package sql_audit

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
const Doc = `SQL审核：表update/delete操作没有指定要具体列，会导致全表变更，请核对`

var Analyzer = &analysis.Analyzer{
	Name:     "xorm_sql_audit",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	//Requires: []*analysis.Analyzer{buildssa.Analyzer},
	Run: run,
}

var expFunc = map[string]bool{
	"In":      true,
	"NotIn":   true,
	"Where":   true,
	"And":     true,
	"Or":      true,
	"ID":      true,
	"Id":      true,
	"Cols":    true,
	"SetExpr": true,
}

func checkCallObj(pass *analysis.Pass, call *ast.CallExpr, checkPreceded bool) (hasCol bool) {
	//check
	if selExpr, ok := call.Fun.(*ast.SelectorExpr); ok {
		funcName := selExpr.Sel.Name

		if true == expFunc[funcName] {
			return true
		}

		if _callExp, ok2 := selExpr.X.(*ast.CallExpr); ok2 {
			hasCol = checkCallObj(pass, _callExp, checkPreceded)
			if hasCol {
				return
			}
		}
	}

	if !checkPreceded {
		return
	}

	// 查看AssignStmt中有没有相关赋值
	caller := findRootCaller(call)
	if nil == caller {
		return
	}
	ident, ok := caller.(*ast.Ident)
	if !ok {
		return
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
	}

	inspect.Nodes(nodeFilter, func(n ast.Node, push bool) (proceed bool) {
		if hasCol {
			proceed = false
			return
		}
		proceed = true
		assignStmt := n.(*ast.AssignStmt)
		if 1 != len(assignStmt.Lhs) {
			return
		}
		assignIdent, ok2 := assignStmt.Lhs[0].(*ast.Ident)
		if !ok2 {
			return
		}
		if assignIdent.Obj != ident.Obj {
			return
		}
		proceed = true // found one node, but still found all
		if 1 != len(assignStmt.Rhs) {
			return
		}
		rhsExpr := assignStmt.Rhs[0]
		if rhsCallExpr, ok3 := rhsExpr.(*ast.CallExpr); ok3 {
			hasCol = checkCallObj(pass, rhsCallExpr, false)
		}
		if hasCol {
			// 只要找到有列条件，立刻停止查询
			proceed = false
		}
		return
	})
	return
}

var ops = map[string]int{"Update": 1, "Delete": 1}

func findRootCaller(c *ast.CallExpr) ast.Expr {
	if s, _ok := c.Fun.(*ast.SelectorExpr); _ok {
		switch _x := s.X.(type) {
		case *ast.CallExpr:
			return findRootCaller(_x)
		case *ast.Ident:
			return s.X
		}
	}
	return nil
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
			if ops[selectExpr.Sel.Name] == 0 {
				return
			}
			if len(callExpr.Args) != ops[selectExpr.Sel.Name] {
				// 限定参数个数
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
				case *types.Named:
					if struct2, ok3 := x.Underlying().(*types.Struct); ok3 {
						procStruct(struct2)
					}
				}
			}
			return
		}

		isXormFunc2 := func(callExpr *ast.CallExpr) (r bool) {
			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				selection := pass.TypesInfo.Selections[selExpr]
				if nil == selection {
					return
				}
				if strings.Contains(selection.Obj().Pkg().Path(), "xorm") {
					r = true
					return
				}
			}
			return
		}

		//isCalledByXorm := func(callExpr *ast.CallExpr) (r bool) {
		//	rootCaller := findRootCaller(callExpr)
		//	if nil == rootCaller {
		//		return false
		//	}
		//
		//	tav := pass.TypesInfo.Types[rootCaller]
		//	if pointer, ok2 := tav.Type.(*types.Pointer); ok2 {
		//		if named, ok3 := pointer.Elem().(*types.Named); ok3 {
		//			if strings.Contains(named.Obj().Pkg().String(), "xorm") {
		//				return true
		//			}
		//		}
		//	}
		//	return false
		//}

		// 根据参数是否含xorm tag判断是否xorm函数
		if r := isXormFunc(callExpr); !r {
			return
		}

		//// 根据根调用者判断是否xorm包中的
		//if r := isCalledByXorm(callExpr); !r {
		//	return
		//}

		// 根据是否xorm包函数判断
		if r2 := isXormFunc2(callExpr); !r2 {
			return
		}

		if !checkCallObj(pass, callExpr, true) {
			common.Reportf(pass, "Ziipin-Garra-XORM-Sql-Audit", callExpr.Pos(),
				fmt.Sprintf("SQL审核：表update/delete操作没有指定要具体列，会导致全表变更，请核对"))
		}

	})

	return nil, nil
}
