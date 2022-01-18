package rules

import (
	"testing"

	"github.com/quasilyte/go-ruleguard/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/delivery-club/delivery-club-rules/pkg"
)

var _ pkg.SQLDb = nil

func TestRules(t *testing.T) {
	testdata := analysistest.TestData()
	if err := analyzer.Analyzer.Flags.Set("rules", "rules.go"); err != nil {
		t.Fatalf("set rules flag: %v", err)
	}
	analysistest.Run(t, testdata, analyzer.Analyzer, "./...")
}
