package main

import (
	"github.com/carsonfeng/garra/analysis/passes/nilcheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		nilcheck.Analyzer,
	)
}
