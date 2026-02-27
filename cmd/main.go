package main

import (
	"github.com/SafroO-Ov/testTaskSelectel/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
