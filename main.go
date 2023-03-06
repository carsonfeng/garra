package main

import (
	"github.com/carsonfeng/garra/analysis/passes/sawa/global_vars_modified"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		//nilcheck.Analyzer,
		//sawago.Analyzer,
		//xorm_index_type_mismatch.Analyzer,
		//xorm_sql_audit.Analyzer,
		//for_get_data.Analyzer,
		global_vars_modified.Analyzer,
	)
}
