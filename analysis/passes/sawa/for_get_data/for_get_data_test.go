// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package for_get_data_test

import (
	"github.com/carsonfeng/garra/analysis/passes/sawa/for_get_data"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestA(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), for_get_data.Analyzer, "a")
}
