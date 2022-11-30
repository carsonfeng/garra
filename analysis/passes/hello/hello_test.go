// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hello_test

import (
	"fmt"
	"github.com/carsonfeng/garra/analysis/passes/hello"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func Test(t *testing.T) {
	result := analysistest.Run(t, analysistest.TestData(), hello.Analyzer)
	for i, r := range result {
		fmt.Printf("result %d: %+v", i, *r)
	}
}
