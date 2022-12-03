package main

import (
	"github.com/carsonfeng/garra/analysis/passes/nilcheck"
	"github.com/carsonfeng/garra/analysis/passes/sawago"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		nilcheck.Analyzer,
		sawago.Analyzer,
	)
}
