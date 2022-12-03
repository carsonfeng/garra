// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nilcheck_test

import (
	"github.com/carsonfeng/garra/analysis/passes/nilcheck"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestA(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nilcheck.Analyzer, "a")
}

func TestB(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nilcheck.Analyzer, "b")
}

func TestC(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nilcheck.Analyzer, "c")
}

func TestD(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nilcheck.Analyzer, "d")
}

func TestE(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), nilcheck.Analyzer, "e")
}
