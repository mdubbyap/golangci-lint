package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"github.com/golangci/golangci-lint/pkg/lint/linter"
	"github.com/mdubbyap/tedo-lint/tedocheck"
	"golang.org/x/tools/go/analysis"
)

const tedoName = "tedo"

func NewTedocheck() *goanalysis.Linter {
	analyzer := &analysis.Analyzer{
		Name: goanalysis.TheOnlyAnalyzerName,
		Doc:  goanalysis.TheOnlyanalyzerDoc,
	}
	return goanalysis.NewLinter(
		tedoName,
		"Tool for detection of tedo things",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithContextSetter(func(lintCtx *linter.Context) {
		analyzer.Run = tedocheck.Run
	})
}
