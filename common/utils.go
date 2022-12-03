package common

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/analysis"
)

func Reportf(pass *analysis.Pass, category string, pos token.Pos, format string, args ...interface{}) {
	message := fmt.Sprintf("[Ziipin-Best-Practices] %s [Garra Ver %v]", fmt.Sprintf(format, args...), Version)
	pass.Report(analysis.Diagnostic{
		Pos:      pos,
		Category: category,
		Message:  message,
	})
}
