// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package async_third_call

import (
	"fmt"
	"github.com/carsonfeng/garra/common"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const DEBUG = false
const Doc = `[Ziipin-Best-Practices] 主业务逻辑不影响的第三方操作，能异步的尽量异步`

var Analyzer = &analysis.Analyzer{
	Name: "async_third_call",
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

var targetFunc = map[string]string{
	"GetUser":      "svc.User().GetUserMapCache",
	"GetUserCache": "svc.User().GetUserMapCache",
}

func runFunc(pass *analysis.Pass, fn *ssa.Function) {
	visit := func(blockIndex int, b *ssa.BasicBlock) {
		if b.Comment != "for.body" && b.Comment != "rangeindex.body" {
			return
		}
		for _, ins := range b.Instrs {
			if call, ok := ins.(*ssa.Call); ok {
				callName := call.Call.Value.Name()
				if resolve, found := targetFunc[callName]; found {
					common.Reportf(pass, "Ziipin-Garra-for-get-data", ins.Pos(), fmt.Sprintf("循环批量获取数据(函数%s)时，建议优先使用批量读缓存或者读表的方式(参考：%s)，避免循环单个读", callName, resolve))
				}
			}
		}
	}

	for i, b := range fn.Blocks {
		visit(i, b)
	}
}
