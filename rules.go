package rules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func unusedFormatting(m dsl.Matcher) {
	m.Import("github.com/pkg/errors")

	m.Match(`fmt.Sprintf($f)`,
		`fmt.Printf($f)`,
		`fmt.Fscanf($_, $f)`,
		`fmt.Fprintf($_, $f)`,
		`fmt.Scanf($f)`,
		`fmt.Sscanf($_, $f)`,
		`errors.WithMessagef($_, $f)`,
		`errors.Wrapf($_, $f)`,
	).
		Where(m["f"].Text.Matches(".*")).
		Report(`use function alternative without formatting`)

	m.Match(`fmt.Errorf($f)`,
		`errors.Errorf($f)`,
	).
		Where(m["f"].Text.Matches(".*")).
		Suggest("errors.New($f)").
		Report("use errors.New instead")
}

// from https://github.com/quasilyte/uber-rules
func uncheckedTypeAssert(m dsl.Matcher) {
	m.Match(
		`$_ := $_.($_)`,
		`$_ = $_.($_)`,
		`$_($*_, $_.($_), $*_)`,
		`$_{$*_, $_.($_), $*_}`,
		`$_{$*_, $_: $_.($_), $*_}`).
		Report(`avoid unchecked type assertions as they can panic`)
}

// from https://github.com/quasilyte/go-ruleguard
func rangeCopyVal(m dsl.Matcher) {
	m.Match(`for $_, $x := range $xs { $*_ }`, `for $_, $x = range $xs { $*_ }`).
		Where((m["xs"].Type.Is("[]$_") || m["xs"].Type.Is("[$_]$_")) && m["x"].Type.Size >= 256).
		Report("each iteration copies more than 256 bytes (consider pointers or indexing)").
		At(m["x"])
}

// from https://github.com/quasilyte/go-ruleguard
func rangeExprCopy(m dsl.Matcher) {
	m.Match(`for $_, $_ := range $x { $*_ }`,
		`for $_, $_ = range $x { $*_ }`).
		Where(m["x"].Type.Is("[$_]$_") && m["x"].Type.Size >= 256).
		Report(`copy of $x can be avoided with &$x`).
		At(m["x"]).
		Suggest(`&$x`)
}

func camelCaseNaming(m dsl.Matcher) {
	m.Match(
		`func $x($*_) $*_ { $*_ }`,
		`const $x = $_`, `const $x $_ = $_`,
		`const ($x = $_; $*_)`,          // workaround for https://github.com/quasilyte/go-ruleguard/issues/160
		`const ($_ = $_; $x = $_; $*_)`, // wip: not working yet because previous rule
		`const ($x $_= $_; $*_)`,
		`const ($_ $_ = $_; $x $_= $_; $*_)`, // wip: not working yet because previous rule
		`type $x $_`,
		`$x := $_`,
		`var $x = $_`,
		`var $x $_ = $_`,
	).
		Where(!m["x"].Text.Matches(`^_$`) && (m["x"].Text.Matches(`-`) || m["x"].Text.Matches(`_`))).
		Report("use camelCase naming strategy").
		At(m["x"])
}

func notInformativePackageNaming(m dsl.Matcher) {
	m.Match(`package $x`).
		Where(
			m["x"].Text.Matches(`(^c|C|_(c|C))ommon([A-Z]|_|$)`) ||
				m["x"].Text.Matches(`(^l|L|_(l|L))ib([A-Z]|_|$)`) ||
				m["x"].Text.Matches(`(^u|U|_(u|U))til([A-Z]|_|$)`) ||
				m["x"].Text.Matches(`(^s|S|_(s|S))hared([A-Z]|_|$)`),
		).
		Report("don't use general names to package naming").
		At(m["x"])
}
