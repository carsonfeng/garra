// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nilcheck_test

import (
	"fmt"
	"github.com/carsonfeng/garra/analysis/passes/nilcheck"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func Test(t *testing.T) {
	result := analysistest.Run(t, analysistest.TestData(), nilcheck.Analyzer, "a")
	for i, r := range result {
		fmt.Printf("result %d: %+v", i, *r)
	}
}
