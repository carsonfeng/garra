// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The vet command is a static checker for Go programs. It has pluggable
// analyzers defined using the golang.org/x/tools/go/analysis API, and
// using the golang.org/x/tools/go/packages API to load packages in any
// build system.
//
// Each analyzer flag name is preceded by the analyzer name: -NAME.flag.
// In addition, the -NAME flag itself controls whether the
// diagnostics of that analyzer are displayed. (A disabled analyzer may yet
// be run if it is required by some other analyzer that is enabled.)
package main

import (
	"github.com/carsonfeng/garra/analysis/passes/nilcheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	// This suite of analyzers is applied to all code
	// in GOROOT by GOROOT/src/cmd/vet/all. When adding
	// a new analyzer, update the whitelist used by vet/all,
	// or change its vet command to disable the new analyzer.
	multichecker.Main(
		nilcheck.Analyzer,
	)
}
