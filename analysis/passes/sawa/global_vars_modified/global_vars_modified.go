// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package global_vars_modified

import (
	"fmt"
	"github.com/carsonfeng/garra/common"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
	"strings"
)

const DEBUG = false
const Doc = `[Ziipin-Best-Practices] 全局变量在函数中被修改，确认是否有风险`

var Analyzer = &analysis.Analyzer{
	Name: "global_vars_modified",
	Doc:  Doc,
	//Requires: []*analysis.Analyzer{inspect.Analyzer},
	Requires: []*analysis.Analyzer{buildssa.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ssainput := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	for _, fn := range ssainput.SrcFuncs {
		runFunc(pass, fn)
	}
	return nil, nil
}

func runFunc(pass *analysis.Pass, fn *ssa.Function) {
	if strings.HasPrefix(fn.Name(), "init") {
		return
	}
	visit := func(blockIndex int, b *ssa.BasicBlock) {
		globalAlloc := map[*ssa.Alloc]*ssa.Global{}
		for _, ins := range b.Instrs {
			if store, ok := ins.(*ssa.Store); ok {
				if fieldAddr, ok2 := store.Addr.(*ssa.FieldAddr); ok2 {
					if global, ok3 := fieldAddr.X.(*ssa.Global); ok3 {
						common.Reportf(pass, "Ziipin-Garra-Global-Vars-Modified", ins.Pos(), fmt.Sprintf("全局变量%s字段%d在函数中被修改，二次确认是否有风险", global.Name(), fieldAddr.Field))
					}
					if alloc, ok3 := fieldAddr.X.(*ssa.Alloc); ok3 {
						if globalAlloc[alloc] != nil {
							common.Reportf(pass, "Ziipin-Garra-Global-Vars-Modified", ins.Pos(), fmt.Sprintf("全局变量%s字段%d在函数中被间接修改，二次确认是否有风险", globalAlloc[alloc].Name(), fieldAddr.Field))
						}
					}
				}
				//if alloc, ok4 := store.Addr.(*ssa.Alloc); ok4 {
				//	if unOp, ok5 := store.Val.(*ssa.UnOp); ok5 {
				//		// 一次传递
				//		if global, ok6 := unOp.X.(*ssa.Global); ok6 {
				//			globalAlloc[alloc] = global
				//		}
				//		// 二次传递
				//		if al2, ok6 := unOp.X.(*ssa.Alloc); ok6 {
				//			if globalAlloc[al2] != nil {
				//				globalAlloc[alloc] = globalAlloc[al2]
				//			}
				//		}
				//	}
				//}
			}
		}
	}

	for i, b := range fn.Blocks {
		visit(i, b)
	}
}
