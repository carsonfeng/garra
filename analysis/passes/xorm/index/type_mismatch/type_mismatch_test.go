// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm_index_type_mismatch_test

import (
	xorm_index_type_mismatch "github.com/carsonfeng/garra/analysis/passes/xorm/index/type_mismatch"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestA(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), xorm_index_type_mismatch.Analyzer, "a")
}
