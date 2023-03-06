// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package global_vars_modified_test

import (
	"github.com/carsonfeng/garra/analysis/passes/sawa/global_vars_modified"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestA(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), global_vars_modified.Analyzer, "a")
}
