package main

import (
	"fmt"
	"go/token"
	"os"
	"strings"

	"github.com/quasilyte/go-ruleguard/ruleguard"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/delivery-club/delivery-club-rules/cmd/precompile/rulesdata"
)

//go:generate go run ./precompile/precompile.go -rules ../rules.go -o ./precompile/rulesdata/rulesdata.go

var (
	flagTag     string
	flagDebug   string
	flagDisable string
	flagEnable  string
)

// Version contains extra version info.
// It's initialized via ldflags -X when deliveryClubRules is built with Make.
// Can contain a git hash (dev builds) or a version tag (release builds).
var Version string

func docString() string {
	doc := "execute dynamic gogrep-based rules"
	if Version == "" {
		return doc
	}
	return doc + " (" + Version + ")"
}

// Analyzer exports deliveryClubRules as a analysis-compatible object.
var Analyzer = &analysis.Analyzer{
	Name: "deliveryClubRules",
	Doc:  docString(),
	Run:  runAnalyzer,
}

var globalEngine *ruleguard.Engine

func init() {
	//Analyzer.Flags.StringVar(&flagTag, "tags", "", "comma separated rules tags")
	Analyzer.Flags.StringVar(&flagDisable, "disabled", "", "comma-separated list of enabled groups or skip empty to enable everything")
	Analyzer.Flags.StringVar(&flagEnable, "enabled", "<all>", "comma-separated list of disabled groups or skip empty to enable everything")

	enabledGroups := make(map[string]bool)
	disabledGroups := make(map[string]bool)
	for _, g := range strings.Split(flagDisable, ",") {
		g = strings.TrimSpace(g)
		disabledGroups[g] = true
	}
	if flagEnable != "<all>" {
		for _, g := range strings.Split(flagEnable, ",") {
			g = strings.TrimSpace(g)
			enabledGroups[g] = true
		}
	}

	// TODO add filter by tags after https://github.com/quasilyte/go-ruleguard/pull/376
	loadContext := &ruleguard.LoadContext{
		Fset:       token.NewFileSet(),
		DebugPrint: debugPrint,
		GroupFilter: func(g string) bool {
			whyDisabled := ""
			enabled := flagEnable == "<all>" || enabledGroups[g]
			switch {
			case !enabled:
				whyDisabled = "not enabled by -enabled flag"
			case disabledGroups[g]:
				whyDisabled = "disabled by -disable flag"
			}
			if flagDebug != "" {
				if whyDisabled != "" {
					debugPrint(fmt.Sprintf("(-) %s is %s", g, whyDisabled))
				} else {
					debugPrint(fmt.Sprintf("(+) %s is enabled", g))
				}
			}
			return whyDisabled == ""
		},
	}

	globalEngine = ruleguard.NewEngine()
	globalEngine.InferBuildContext()

	if err := globalEngine.LoadFromIR(loadContext, "rulesdata.go", rulesdata.PrecompiledRules); err != nil {
		fmt.Println("on load ir rules: ", err)
		os.Exit(1)
	}
}

func main() {
	singlechecker.Main(Analyzer)
}

func runAnalyzer(pass *analysis.Pass) (interface{}, error) {
	ctx := &ruleguard.RunContext{
		Debug:      flagDebug,
		DebugPrint: debugPrint,
		Pkg:        pass.Pkg,
		Types:      pass.TypesInfo,
		Sizes:      pass.TypesSizes,
		Fset:       pass.Fset,
		Report: func(data *ruleguard.ReportData) {
			fullMessage := data.Message
			diag := analysis.Diagnostic{
				Pos:     data.Node.Pos(),
				Message: fullMessage,
			}
			if data.Suggestion != nil {
				s := data.Suggestion
				diag.SuggestedFixes = []analysis.SuggestedFix{
					{
						Message: "suggested replacement",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     s.From,
								End:     s.To,
								NewText: s.Replacement,
							},
						},
					},
				}
			}
			pass.Report(diag)
		},
	}

	for _, f := range pass.Files {
		if err := globalEngine.Run(ctx, f); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func debugPrint(s string) {
	fmt.Println("debug:", s)
}
