package main

import (
	"github.com/carsonfeng/garra/analysis/passes/nilcheck"
	"github.com/carsonfeng/garra/analysis/passes/sawa/for_get_data"
	"github.com/carsonfeng/garra/analysis/passes/sawa/global_vars_modified"
	"github.com/carsonfeng/garra/analysis/passes/sawago"
	xorm_index_type_mismatch "github.com/carsonfeng/garra/analysis/passes/xorm/index/type_mismatch"
	xorm_sql_audit "github.com/carsonfeng/garra/analysis/passes/xorm/sql_audit"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		nilcheck.Analyzer,
		sawago.Analyzer,
		xorm_index_type_mismatch.Analyzer,
		xorm_sql_audit.Analyzer,
		for_get_data.Analyzer,
		global_vars_modified.Analyzer,
	)
}
