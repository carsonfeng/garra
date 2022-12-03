// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sawago defines an Analyzer that checks if nil object is used
package sawago

import (
	"fmt"
	"github.com/carsonfeng/garra/common"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const DEBUG = false
const Doc = `[Ziipin-Best-Practices] 如无特殊情况，请用services.AsynHandle方法来调用，不要自己直接go [Sawa]

example:

Suggest:
func (dao *UserService)testFunc(){
	asynHandle(func(svc *Svc) {
		XXXX
	})
}

NOT SUGGEST:
func (dao *UserService)testFunc(){
	go func() {
	}
}

`

var Analyzer = &analysis.Analyzer{
	Name: "sawago",
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
	pkgName := fn.Package().Pkg.Name()
	if "daos" != pkgName && "services" != pkgName && "views" != pkgName {
		return
	}
	fnName := fn.Name()
	if "AsynHandle" == fnName || "AsyncHandle" == fnName {
		return
	}
	for _, b := range fn.Blocks {
		for _, instr := range b.Instrs {
			if _, ok := instr.(*ssa.Go); ok {
				common.Reportf(pass, "Ziipin-Sawa-Go", instr.Pos(), fmt.Sprintf("如无特殊情况，协程请用services.AsynHandle或daos.AsynHandle方法来调用，不要自己直接go"))
			}
		}
	}
}
