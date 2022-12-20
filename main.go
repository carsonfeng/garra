package main

import (
	"github.com/carsonfeng/garra/analysis/passes/nilcheck"
	xorm_index_type_mismatch "github.com/carsonfeng/garra/analysis/passes/xorm/index/type_mismatch"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		nilcheck.Analyzer,
		//sawago.Analyzer,
		xorm_index_type_mismatch.Analyzer,
	)
}

func useRids(args ...interface{}) {

}

func mainmain2() {
	rids := make([]string, 0, 10)
	useRids(rids)
}
