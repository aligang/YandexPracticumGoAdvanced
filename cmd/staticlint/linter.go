package main

import (
	"github.com/gostaticanalysis/nilerr"
	"github.com/gostaticanalysis/sqlrows/passes/sqlrows"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
	"strings"
)

// main Linting can be started as go run . ../cmd/...
func main() {
	checkList := []*analysis.Analyzer{}

	standartChecks := []*analysis.Analyzer{
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
	}
	checkList = append(checkList, standartChecks...)

	thirdPartyChecks := []*analysis.Analyzer{
		sqlrows.Analyzer,
		nilerr.Analyzer,
	}
	checkList = append(checkList, thirdPartyChecks...)

	customStaticCheckFlags := map[string]bool{
		"-ST1000": true,
		"-ST1003": true,
		"-ST1016": true,
	}

	for _, v := range staticcheck.Analyzers {
		// добавляем в массив нужные проверки
		if strings.HasPrefix(v.Analyzer.Name, "SA") || customStaticCheckFlags[v.Analyzer.Name] {
			checkList = append(checkList, v.Analyzer)
		}
	}
	checkList = append(checkList, MainExitAnalyzer)
	multichecker.Main(checkList...)
}
