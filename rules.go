package rules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func unusedFormatting(m dsl.Matcher) {
	m.Import("github.com/pkg/errors")

	m.Match(`fmt.Sprintf($_)`, `errors.WithMessagef($_, $_)`, `errors.Wrapf($_, $_)`, `errors.Errorf($_)`).
		Report(`use function alternative without formatting`)
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

// from https://github.com/quasilyte/uber-rules
func ifacePtr(m dsl.Matcher) {
	m.Match(`*$x`).
		Where(m["x"].Type.Underlying().Is(`interface{ $*_ }`)).
		Report(`don't use pointers to an interface`)
}

func camelCaseNaming(m dsl.Matcher) {
	m.Match(
		`func $x($*_) $*_ { $*_ }`,
		`func ($_) $x($*_) $*_ { $*_ }`,
		`func ($_ $_) $x($*_) $*_ { $*_ }`,
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
			m["x"].Text.Matches(`(^c|C|_(c|C))ommon([A-Z]|_|$|\d)`) ||
				m["x"].Text.Matches(`(^l|L|_(l|L))ib([A-Z]|_|$|\d)`) ||
				m["x"].Text.Matches(`(^u|U|_(u|U))til([A-Z]|_|$|\d)`) ||
				m["x"].Text.Matches(`(^s|S|_(s|S))hared([A-Z]|_|$|\d)`),
		).
		Report("don't use general names to package naming").
		At(m["x"])
}

func getterNaming(m dsl.Matcher) {
	m.Match(
		`func ($x $_) $name($*_) $*_ { return $x.$_ }`,
		`func ($x $_) $name($*_) $*_ { return $_($x.$_) }`,
		`func ($x $_) $name($*_) $*_ { return $_($x) }`,
		`func ($x $_) $name($*_) $*_ { return $_(*$x) }`,
	).
		Where(m["name"].Text.Matches(`(^g|G|_(g|G))et([A-Z]|$|_|\d)`)).
		Report(`don't use 'get' in getter functions`).
		At(m["name"])
}

func oneMethodInterfaceNaming(m dsl.Matcher) {
	m.Match(`type $name interface{ $method ($*_) $*_ }`).
		Where(!m["name"].Text.Matches(`\wer$`)).
		Report("change interface name to $method + 'er' pattern").
		At(m["name"])
}

func interfaceWordInInterfaceDeclaration(m dsl.Matcher) {
	m.Match(`type $name interface{ $*_ }`).
		Where(m["name"].Text.Matches(`(^i|I|_(i|I))nterface([A-Z]|_|$|\d)`)).
		Report(`don't use 'interface' word' in interface declaration'`).
		At(m["name"])
}

func simplifyErrorReturn(m dsl.Matcher) {
	m.Match(`if $*_, $err = $f($*args); $err != nil { return $err }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return $err }; return nil`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`).
		Suggest(`$*_, err := $f($args); return err`).
		At(m["f"])
}

func simplifyErrorReturnWithErrorsPkg(m dsl.Matcher) {
	m.Import("github.com/pkg/errors")

	m.Match(
		`if $*_, $err = $f($*args); $err != nil { return errors.WithMessagef($err, $*methodArgs) }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return errors.WithMessagef($err, $*methodArgs) }; return nil`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`).
		Suggest(`$*_, err := $f($args); return errors.WithMessagef($err, $*methodArgs)`).
		At(m["f"])

	m.Match(
		`if $*_, $err = $f($*args); $err != nil { return errors.WithMessage($err, $*methodArgs) }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return errors.WithMessage($err, $*methodArgs) }; return nil`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`).
		Suggest(`$*_, err := $f($args); return errors.WithMessage($err, $*methodArgs)`).
		At(m["f"])

	m.Match(
		`if $*_, $err = $f($*args); $err != nil { return errors.Wrap($err, $*methodArgs) }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return errors.Wrap($err, $*methodArgs) }; return nil`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`).
		Suggest(`$*_, err := $f($args); return errors.Wrap($err, $*methodArgs)`).
		At(m["f"])

	m.Match(
		`if $*_, $err = $f($*args); $err != nil { return errors.Wrapf($err, $*methodArgs) }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return errors.Wrapf($err, $*methodArgs) }; return nil`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`).
		Suggest(`$*_, err := $f($args); return errors.Wrapf($err, $*methodArgs)`).
		At(m["f"])

	m.Match(
		`if $*_, $err = $f($*args); $err != nil { return errors.WithStack($err, $*methodArgs) }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return errors.WithStack($err, $*methodArgs) }; return nil`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`).
		Suggest(`$*_, err := $f($args); return errors.WithStack($err, $*methodArgs)`).
		At(m["f"])
}
