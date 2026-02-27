package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "testTaskSelectel",
	Doc:  "linter for checking logs",
	Run:  run,
}
