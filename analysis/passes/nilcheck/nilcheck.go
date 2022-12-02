// Copyright 2022 Ziipin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package nilcheck defines an Analyzer that checks if nil object is used
package nilcheck

import (
	"fmt"
	"github.com/carsonfeng/garra/common"
	"go/constant"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const DEBUG = true
const Doc = `[Ziipin-Best-Practices] check if object is used with not-nil error


For example:
	user, err := getUser();
	if nil != err{
		niuhe.LogError("some error")
	}
	uid := user.GetUid() // <-- may cause nil pointer panic
`

var Analyzer = &analysis.Analyzer{
	Name: "nilcheck",
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
	//fmt.Printf("parsing func: %s\n", fn.Name())
	//fmt.Printf("fn: %#v\n", fn)
	//return
	if DEBUG && "NilObjFunc" != fn.Name() {
		return
	}
	reportf := func(category string, pos token.Pos, format string, args ...interface{}) {
		pass.Report(analysis.Diagnostic{
			Pos:      pos,
			Category: category,
			Message:  fmt.Sprintf(format, args...),
		})
	}

	var visit func(blockIndex int, b *ssa.BasicBlock)

	hasReturnOrPanic := func(b *ssa.BasicBlock) (bool, error) {
		for _, ins := range b.Instrs {
			if _, ok := ins.(*ssa.Return); ok {
				return true, nil
			}

			if _, ok := ins.(*ssa.Panic); ok {
				return true, nil
			}
		}
		return false, nil
	}

	isNilValue := func(v ssa.Value) bool {
		if con, ok := v.(*ssa.Const); ok {
			return con.Value == constant.Value(nil)
		}
		return false
	}

	isErrorType := func(v ssa.Value) bool {
		t := v.Type().String()
		return t == "error"
	}

	findLastCallExtracts := func(in []ssa.Instruction) (m map[*ssa.Extract]bool) {
		m = map[*ssa.Extract]bool{}
		for i := len(in) - 1; i >= 0; i-- {
			if call, ok := in[i].(*ssa.Call); ok {
				ref := call.Referrers()
				if nil == ref {
					return
				}
				for _, r := range *ref {
					if ext, ok2 := r.(*ssa.Extract); ok2 && "error" != ext.Type().String() {
						m[ext] = true
					}
				}
				return
			}
		}
		return
	}

	checkIfDone := func(originCallBlock *ssa.BasicBlock, b *ssa.BasicBlock) {
		// check if this block use unsafe obj
		//fmt.Printf("check if this block[%d] use unsafe obj\n", blockIndex)
		for _, instr2 := range b.Instrs {
			//fmt.Printf("instr2: %#v\n", instr2)
			if call, ok := instr2.(*ssa.Call); ok {
				//fmt.Printf("call args: %#v\n", call.Call.Args)
				if "len" == call.Call.Value.Name() {
					continue
				}

				m := findLastCallExtracts(originCallBlock.Instrs)
				for _, arg := range call.Call.Args {
					if ext, ok3 := arg.(*ssa.Extract); ok3 {
						if m[ext] {
							reportf("Ziipin-Garra-nilcheck", call.Pos(), fmt.Sprintf("[Ziipin-Best-Practices] call object's method/field with non-nil error will always panic. [Garra ver: %s]", common.Version))
						}
					}
				}
			}
		}
	}

	checkIfThen := func(originCallBlock *ssa.BasicBlock, b *ssa.BasicBlock) {
		if has, err := hasReturnOrPanic(b); nil == err && !has {
			if len(b.Succs) > 0 && "if.done" == b.Succs[0].Comment {
				checkIfDone(originCallBlock, b.Succs[0])
			}
		}
	}

	visit = func(blockIndex int, b *ssa.BasicBlock) {
		insLen := len(b.Instrs)
		if 0 == insLen {
			return
		}
		// check if last instruction is if
		instr := b.Instrs[insLen-1]
		//fmt.Printf("parsing block [%d]: %#v\n", blockIndex, *b)
		//for _, t := range b.Instrs {
		//	fmt.Printf("instr: %#v\n", t)
		//}
		if If, ok := instr.(*ssa.If); ok {

			// nil != err
			if binOp, ok2 := If.Cond.(*ssa.BinOp); ok2 && token.NEQ == binOp.Op {
				//fmt.Printf("binOp: %#v \nX: %#v\n Y: %#v\n", *binOp, binOp.X, binOp.Y)
				if (isNilValue(binOp.X) && isErrorType(binOp.Y)) || (isNilValue(binOp.Y) && isErrorType(binOp.X)) {
					m := findLastCallExtracts(b.Instrs)
					if len(m) > 0 {
						if len(If.Block().Succs) > 0 && If.Block().Succs[0].Comment == "if.then" {
							ifThenBlock := If.Block().Succs[0]
							checkIfThen(b, ifThenBlock)
						}
					}
				}
			}
		}

	}

	for i, b := range fn.Blocks {
		visit(i, b)
	}
}
