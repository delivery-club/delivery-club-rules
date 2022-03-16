package rules

import (
	"github.com/quasilyte/go-ruleguard/dsl"

	_ "github.com/delivery-club/delivery-club-rules/pkg"
)

var Bundle = dsl.Bundle{}

//doc:summary   Detects unused formatting functionality
//doc:tags      style
//doc:before    fmt.Sprintf("42")
//doc:after     fmt.Sprint("42")
func unusedFormatting(m dsl.Matcher) {
	m.Match(`fmt.Sprintf($_)`).
		Report(`use function alternative without formatting`)
}

//doc:summary   Detects unused formatting functionality
//doc:tags      style
//doc:before    errors.WithMessagef(err, "on example")
//doc:after     errors.WithMessage(err, "on example")
func unusedErrorsFormatting(m dsl.Matcher) {
	m.Import("github.com/pkg/errors")

	m.Match(`errors.WithMessagef($_, $_)`, `errors.Wrapf($_, $_)`, `errors.Errorf($_)`).
		Report(`use function alternative without formatting`)
}

//doc:summary   Detects unchecked type assertation
//doc:tags      diagnostic
//doc:before    foo := bar.(string) // var bar interface{}
//doc:after     foo, ok := bar.(string)
func uncheckedTypeAssert(m dsl.Matcher) {
	m.Match(
		`$_ := $_.($_)`,
		`$_ = $_.($_)`,
		`$_($*_, $_.($_), $*_)`,
		`$_{$*_, $_.($_), $*_}`,
		`$_{$*_, $_: $_.($_), $*_}`,
		`$_ <- $_.($_)`,
		`$_{$*_, $_.($_): $_, $*_}`).
		Report(`avoid unchecked type assertions as they can panic`)
}

//doc:summary   Detects copy big structs in loop body
//doc:tags      performance
//doc:before    for _, x := range xs { myfunc(x) } // for example var xs [][2048]string
//doc:after     for i := range xs { myfunc(xs[i]) }
func rangeCopyVal(m dsl.Matcher) {
	m.Match(`for $_, $x := range $xs { $*_ }`, `for $_, $x = range $xs { $*_ }`).
		Where((m["xs"].Type.Is("[]$_") || m["xs"].Type.Is("[$_]$_")) && m["x"].Type.Size >= 256).
		Report("each iteration copies more than 256 bytes (consider pointers or indexing)").
		At(m["x"])
}

//doc:summary   Detects big array copy in loop
//doc:tags      performance
//doc:before    for _, x := range xs { myfunc(x) } // var xs [2048]string
//doc:after     for _, x := range &xs { myfunc(x) }
func rangeExprCopy(m dsl.Matcher) {
	m.Match(`for $_, $_ := range $x { $*_ }`,
		`for $_, $_ = range $x { $*_ }`).
		Where(m["x"].Type.Is("[$_]$_") && m["x"].Type.Size >= 256).
		Report(`copy of $x can be avoided with &$x`).
		At(m["x"]).
		Suggest(`&$x`)
}

//doc:summary   Detects pointer to interface{}
//doc:tags      performance
func ifacePtr(m dsl.Matcher) {
	m.Match(`*$x`).
		Where(m["x"].Type.Underlying().Is(`interface{ $*_ }`)).
		Report(`don't use pointers to an interface`)
}

//doc:summary   Detects non camel case naming strategy for functions, constants, types, variables
//doc:tags      style
//doc:before    type my_type struct {}
//doc:after     type myType struct {}
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

//doc:summary   Detects general names for package
//doc:tags      style
//doc:before    package lib
//doc:after     package concreteLib
func notInformativePackageNaming(m dsl.Matcher) {
	m.Match(`package $x`).
		Where(m["x"].Text.Matches(`(^c|C|_(c|C))ommon([A-Z]|_|$|\d)`) ||
			m["x"].Text.Matches(`(^l|L|_(l|L))ib([A-Z]|_|$|\d)`) ||
			m["x"].Text.Matches(`(^u|U|_(u|U))til([A-Z]|_|$|\d)`) ||
			m["x"].Text.Matches(`(^s|S|_(s|S))hared([A-Z]|_|$|\d)`)).
		Report("don't use general names to package naming").
		At(m["x"])
}

//doc:summary   Detects 'get' word in getter functions
//doc:tags      style
//doc:before    func (x myType) getValue() string { return x.v }
//doc:after     func (x myType) value() string { return x.v }
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

// disabled until https://github.com/go-critic/go-critic/issues/1176
//func oneMethodInterfaceNaming(m dsl.Matcher) {
//	m.Match(`type $name interface{ $method ($*_) $*_ }`).
//		Where(m["name"].Text.Matches(`^[A-Z]`) && !m["name"].Text.Matches(`\w(er|or|ar)$`)).
//		Report("change interface name to $method + 'er|or|ar' pattern").
//		At(m["name"])
//}

//doc:summary   Detects 'interface' word in interface declarations
//doc:tags      style
//doc:before    type interfaceDb interface { }
//doc:after     type db interface { }
func interfaceWordInInterfaceDeclaration(m dsl.Matcher) {
	m.Match(`type $name interface{ $*_ }`).
		Where(m["name"].Text.Matches(`(^i|I|_(i|I))nterface([A-Z]|_|$|\d)`)).
		Report(`don't use 'interface' word' in interface declaration'`).
		At(m["name"])
}

// TODO add conditional suggest after implement https://github.com/quasilyte/go-ruleguard/issues/301
//doc:summary   Detects expressions that can be rewritten in form without 'if' usage
//doc:tags      style
//doc:before    err := myFunc(); if err != nil { return err }; return nil
//doc:after     return myFunc()
func simplifyErrorReturn(m dsl.Matcher) {
	m.Match(`if $*_, $err = $f($*args); $err != nil { return $err }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return $err }; return nil`,
		`$*_, $err = $f($*args); if $err != nil { return $err }; return nil`,
		`var $*_, $err = $f($*args); if $err != nil { return $err }; return nil`,
		`$*_, $err := $f($*args); if $err != nil { return $err }; return nil`,
		`if $*_, $err = $f($*args); $err != nil { return $err }; return $err`,
		`if $*_, $err := $f($*args); $err != nil { return $err }; return $err`,
		`$*_, $err = $f($*args); if $err != nil { return $err }; return $err`,
		`var $*_, $err = $f($*args); if $err != nil { return $err }; return $err`,
		`$*_, $err := $f($*args); if $err != nil { return $err }; return $err`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`)
}

//doc:summary   Detects expressions that can be rewritten in form without 'if' usage
//doc:tags      style
//doc:before    err := myFunc(); if err != nil { return errors.WithMessage(err, "on myFunc:") }; return nil
//doc:after     return errors.WithMessage(myFunc(), "on myFunc:")
func simplifyErrorReturnWithErrorsPkg(m dsl.Matcher) {
	m.Import("github.com/pkg/errors")

	m.Match(
		`if $*_, $err = $f($*args); $err != nil { return errors.$_($err, $*methodArgs) }; return nil`,
		`if $*_, $err := $f($*args); $err != nil { return errors.$_($err, $*methodArgs) }; return nil`,
		`$*_, $err = $f($*args); if $err != nil { return errors.$_($err, $*methodArgs) }; return nil`,
		`var $*_, $err = $f($*args); if $err != nil { return errors.$_($err, $*methodArgs) }; return nil`,
		`$*_, $err := $f($*args); if $err != nil { return errors.$_($err, $*methodArgs) }; return nil`,
		`if $*_, $err = $f($*args); $err != nil { return errors.$_($err, $*methodArgs) }; return $err`,
		`if $*_, $err := $f($*args); $err != nil { return errors.$_($err, $*methodArgs) }; return $err`,
		`$*_, $err = $f($*args); if $err != nil { return errors.$_($err, $*methodArgs) }; return $err`,
		`var $*_, $err = $f($*args); if $err != nil { return errors.$_($err, $*methodArgs) }; return $err`,
		`$*_, $err := $f($*args); if $err != nil { return errors.$_($err, $*methodArgs) }; return $err`,
	).
		Where(m["err"].Type.Implements("error")).
		Report(`may be simplified to return error without if statement`)
}

//TODO: too wide for production usage now
//func isBuiltinInterface(ctx *dsl.VarFilterContext) bool {
//	return types.Implements(ctx.Type, ctx.GetInterface("error")) || types.Implements(ctx.Type, ctx.GetInterface("context.Context"))
//}
//
//func returnConcreteInsteadInterface(m dsl.Matcher) {
//	m.Match(`func $name($*_) $arg { $*_ }`,
//		`func ($_ $_) $name($*_) $arg { $*_ }`,
//		`func ($_) $name($*_) $arg { $*_ }`,
//		`func $name($*_) ($arg, $_) { $*_ }`,         //wip: not working yet
//		`func ($_ $_) $name($*_) ($arg, $_) { $*_ }`, //wip: not working yet
//		`func ($_) $name($*_) ($arg, $_) { $*_ }`,    //wip: not working yet
//	).
//		Where(m["name"].Text.Matches(`^[A-Z]`) &&
//			(m["arg"].Type.Underlying().Is(`interface{ $*_ }`) && !m["arg"].Filter(isBuiltinInterface))).
//		Report(`in exported functions return concrete type instead of interface`).
//		At(m["name"])
//}

//doc:summary   Detects loops inside functions that use defer
//doc:tags      diagnostic
//doc:before    for _, filename := range []string{"foo", "bar"} { f, err := os.Open(filename); defer f.Close() }
//doc:after     func process(filename string) {
//	 				f, err := os.Open(filename)
//					defer f.Close()
//				}
//				/* ... */
//				for _, filename := range []string{"foo", "bar"} {
//					process(filename)
//				}
func deferInLoop(m dsl.Matcher) {
	m.Match(
		`for $*_; $*_; $*_ { $*_; defer func($*args) $*_ { $*_ }($*_); $*_ }`,
		`for $*_; $*_; $*_ { $*_; defer $_($*args); $*_ }`,

		`for { $*_; defer func($*args) $*_ { $*_ }($*_); $*_ }`,
		`for { $*_; defer $_($*args); $*_ }`,

		`for $_, $_ := range $_ { $*_; defer func($*args) $*_ { $*_ }($*_); $*_ }`,
		`for $_, $_ := range $_ { $*_; defer $_($*args); $*_ }`,

		`for $_, $_ = range $_ { $*_; defer func($*args) $*_ { $*_ }($*_); $*_ }`,
		`for $_, $_ = range $_ { $*_; defer $_($*args); $*_ }`,

		`for $_ := range $_ { $*_; defer func($*args) $*_ { $*_ }($*_); $*_ }`,
		`for $_ := range $_ { $*_; defer $_($*args); $*_ }`,

		`for $_ = range $_ { $*_; defer func($*args) $*_ { $*_ }($*_); $*_ }`,
		`for $_ = range $_ { $*_; defer $_($*args); $*_ }`,

		`for range $_ { $*_; defer func($*args) $*_ { $*_ }($*_); $*_ }`,
		`for range $_ { $*_; defer $_($*args); $*_ }`,
	).
		Report(`Possible resource leak, 'defer' is called in the 'for' loop`).
		At(m["args"])
}

//doc:summary   Detects queries to database without context
//doc:tags      diagnostic
//doc:before    db.Exec(`SELECT 1`)
//doc:after     db.ExecContext(ctx, `SELECT 1`)
func queryWithoutContext(m dsl.Matcher) {
	// in this rule we check all libraries which extend std sql lib
	// but for check methods that have names different from the std sql,
	// add new ones to match method below

	m.Import(`github.com/delivery-club/delivery-club-rules/pkg`)

	m.Match(
		`$db.Query($*_)`,
		`$db.Queryx($*_)`,
		`$db.QueryRow($*_)`,
		`$db.QueryRowx($*_)`,
		`$db.NamedQuery($*_)`,

		`$db.Exec($*_)`,
		`$db.MustExec($*_)`,
		`$db.NamedExec($*_)`,

		`$db.Get($*_)`,
		`$db.Select($*_)`,

		`$db.Prepare($*_)`,
		`$db.Preparex($*_)`,
		`$db.PrepareNamed($*_)`,

		`$db.Ping($*_)`,
		`$db.Begin($*_)`,
		`$db.MustBegin($*_)`,

		`$db.Stmt($*_)`,
		`$db.Stmtx($*_)`,
		`$db.NamedStmt($*_)`,
	).
		Where(m["db"].Object.Is("Var") &&
			(m["db"].Type.Implements(`pkg.SQLDb`) || m["db"].Type.Implements(`pkg.SQLStmt`) || m["db"].Type.Implements(`pkg.SQLTx`))).
		Report(`don't send query to external storage without context`).
		At(m["db"])
}

//doc:summary   Detects regular expression compilation in loop
//doc:tags      performance
//doc:before    for { if regexp.MatchString("\d", "123") { /*...*/ } }
//doc:after     dig := regexp.MustCompile("\d"); for { if dig.MatchString("123") { /*...*/ } }
func regexpCompileInLoop(m dsl.Matcher) {
	m.Match(
		`for $*_; $*_; $*_ { $*_; $*_ = regexp.$method($s, $*args); $*_ }`,
		`for { $*_; $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ := range $_ { $*_; $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ = range $_ { $*_; $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_ := range $_ { $*_; $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_ = range $_ { $*_; $*_ = regexp.$method($s, $*args); $*_ }`,
		`for range $_ { $*_; $*_ = regexp.$method($s, $*args); $*_ }`,

		`for $*_; $*_; $*_ { $*_; $*_ := regexp.$method($s, $*args); $*_ }`,
		`for { $*_; $*_ := regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ := range $_ { $*_; $*_ := regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ = range $_ { $*_; $*_ := regexp.$method($s, $*args); $*_ }`,
		`for $_ := range $_ { $*_; $*_ := regexp.$method($s, $*args); $*_ }`,
		`for $_ = range $_ { $*_; $*_ := regexp.$method($s, $*args); $*_ }`,
		`for range $_ { $*_; $*_ := regexp.$method($s, $*args); $*_ }`,

		`for $*_; $*_; $*_ { $*_; var $*_ = regexp.$method($s, $*args); $*_ }`,
		`for { $*_; var $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ := range $_ { $*_; var $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ = range $_ { $*_; var $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_ := range $_ { $*_; var $*_ = regexp.$method($s, $*args); $*_ }`,
		`for $_ = range $_ { $*_; var $*_ = regexp.$method($s, $*args); $*_ }`,
		`for range $_ { $*_; var $*_ = regexp.$method($s, $*args); $*_ }`,

		`for $*_; $*_; $*_ { $*_; var $*_ $_ = regexp.$method($s, $*args); $*_ }`,
		`for { $*_; var $*_ $_ = regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ := range $_ { $*_; var $*_ $_ = regexp.$method($s, $*args); $*_ }`,
		`for $_, $_ = range $_ { $*_; var $*_ $_ = regexp.$method($s, $*args); $*_ }`,
		`for $_ := range $_ { $*_; var $*_ $_ = regexp.$method($s, $*args); $*_ }`,
		`for $_ = range $_ { $*_; var $*_ $_ = regexp.$method($s, $*args); $*_ }`,
		`for range $_ { $*_; var $*_ $_ = regexp.$method($s, $*args); $*_ }`,
	).
		At(m["s"]).
		Where(m["s"].Const && m["method"].Text.Matches(`Compile|MustCompilePOSIX|CompilePOSIX|Match|MatchString|MatchReader|MustCompile`)).
		Report(`don't compile regex in the loop, move it outside of the loop`)
}

//doc:summary   Detects unreleased resources
//doc:tags      diagnostic
//doc:before    s, _ := os.Open("foo.txt"); s.Read(body); return body
//doc:after     s, _ := os.Open("foo.txt"); defer s.Close(); s.Read(body); return body
func unclosedResource(m dsl.Matcher) {
	varEscapeFunction := func(x dsl.Var) bool {
		return x.Contains(`$_($*_, $res, $*_)`) || x.Contains(`$_{$*_, $res, $*_}`) ||
			x.Contains(`$_{$*_, $_: $res, $*_}`) || x.Contains(`$_ <- $res`) ||
			x.Contains(`return $*_, $res, $*_`) || x.Contains(`$_[$_] = $res`) ||
			x.Contains(`$_[$res] = $_`) || x.Contains(`$_ = $res;`) || x.Contains(`$_ := $res;`) ||
			x.Contains(`var $_ = $res;`) || x.Contains(`var $_ $_ = $res;`)
	}

	m.Match(`$res, $err := $open($*_); $*body`,
		`$res, $err = $open($*_); $*body`,
		`var $res, $err = $open($*_); $*body`,
	).
		Where(
			m["res"].Type.Implements(`io.Closer`) &&
				!m["res"].Object.IsGlobal() &&
				m["err"].Type.Implements(`error`) &&
				!m["body"].Contains(`$res.Close()`) &&
				!varEscapeFunction(m["body"]),
		).
		Report(`$res.Close() should be deferred right after the $open error check`).
		At(m["res"])
}

//doc:summary   Detects unreleased timer
//doc:tags      performance
//doc:before    timer := time.NewTimer(time.Second); select { case <-timer.C: return nil; default: return nil }
//doc:after     timer := time.NewTimer(time.Second); defer timer.Stop(); select { case <-timer.C: return nil; default: return nil }
func unstoppedTimer(m dsl.Matcher) {
	varEscapeFunction := func(x dsl.Var) bool {
		return x.Contains(`$_($*_, $x, $*_)`) || x.Contains(`$_{$*_, $x, $*_}`) ||
			x.Contains(`$_{$*_, $_: $x, $*_}`) || x.Contains(`$_ <- $x`) ||
			x.Contains(`return $*_, $x, $*_`) || x.Contains(`$_[$_] = $x`) ||
			x.Contains(`$_[$x] = $_`) || x.Contains(`$_ = $x;`) || x.Contains(`$_ := $x;`) ||
			x.Contains(`var $_ = $x;`) || x.Contains(`var $_ $_ = $x;`)
	}

	m.Match(`$x := time.NewTimer($_); $*body`,
		`$x = time.NewTimer($_); $*body`,
		`var $x = time.NewTimer($_); $*body`,
		`var $x $_ = time.NewTimer($_); $*body`).
		Where(!m["x"].Object.IsGlobal() && !m["body"].Contains(`$x.Stop()`) && !varEscapeFunction(m["body"])).
		Report(`unstopped timer`).
		At(m["x"])
}

//doc:summary   Detects unreleased ticker
//doc:tags      performance
//doc:before    ticker := time.NewTicker(time.Second); select { case <-ticker.C: return nil; default: return nil }
//doc:after     ticker := time.NewTicker(time.Second); defer ticker.Stop(); select { case <-ticker.C: return nil; default: return nil }
func unstoppedTicker(m dsl.Matcher) {
	varEscapeFunction := func(x dsl.Var) bool {
		return x.Contains(`$_($*_, $x, $*_)`) || x.Contains(`$_{$*_, $x, $*_}`) ||
			x.Contains(`$_{$*_, $_: $x, $*_}`) || x.Contains(`$_ <- $x`) ||
			x.Contains(`return $*_, $x, $*_`) || x.Contains(`$_[$_] = $x`) ||
			x.Contains(`$_[$x] = $_`) || x.Contains(`$_ = $x;`) || x.Contains(`$_ := $x;`) ||
			x.Contains(`var $_ = $x;`) || x.Contains(`var $_ $_ = $x;`)
	}

	m.Match(`$x := time.NewTicker($_); $*body`,
		`$x = time.NewTicker($_); $*body`,
		`var $x = time.NewTicker($_); $*body`,
		`var $x $_ = time.NewTicker($_); $*body`).
		Where(!m["x"].Object.IsGlobal() && !m["body"].Contains(`$x.Stop()`) && !varEscapeFunction(m["body"])).
		Report(`unstopped ticker`).
		At(m["x"])
}

//doc:summary   Detects expressions that can be rewritten in one 'if' form
//doc:tags      style
//doc:before    err := myFunc(); if err != nil { println(err) }
//doc:after     if err := myFunc(); err != nil { println(err) }
func simplifyErrorCheck(m dsl.Matcher) {
	m.Match(`$err := $f($*args); if $err != nil { $*body }`).
		Where(m["err"].Type.Implements("error") &&
			m["f"].Text.Matches("(?s)^.{0,40}$") && m["args"].Text.Matches("(?s)^.{0,40}$")). /// TODO: check that chars count in line <= 120
		Report(`error check can be simplified in one line`).
		Suggest(`if $err := $f($args); $err != nil { $body }`)

	m.Match(`$err = $f($*args); if $err != nil { $*body }`).
		Where(m["err"].Type.Implements("error") &&
			m["f"].Text.Matches("(?s)^.{0,40}$") && m["args"].Text.Matches("(?s)^.{0,40}$")).
		Report(`error check can be simplified in one line`).
		Suggest(`if $err = $f($args); $err != nil { $body }`)

	m.Match(`var $err = $f($*args); if $err != nil { $*body }`).
		Where(m["err"].Type.Implements("error") &&
			m["f"].Text.Matches("(?s)^.{0,40}$") && m["args"].Text.Matches("(?s)^.{0,40}$")).
		Report(`error check can be simplified in one line`).
		Suggest(`if $err := $f($args); $err != nil { $body }`)
}

//doc:summary   Non-pointer values in sync.Pool involve extra allocation
//doc:tags      performance
func syncPoolNonPtr(m dsl.Matcher) {
	isPointer := func(x dsl.Var) bool {
		return x.Type.Underlying().Is("*$_") || x.Type.Underlying().Is("chan $_") ||
			x.Type.Underlying().Is("map[$_]$_") || x.Type.Underlying().Is("interface{$*_}") ||
			x.Type.Underlying().Is(`func($*_) $*_`) || x.Type.Underlying().Is(`unsafe.Pointer`)
	}

	m.Match(`$x.Put($y)`).
		Where(m["x"].Type.Is("sync.Pool") && !isPointer(m["y"])).
		Report(`non-pointer values in sync.Pool involve extra allocation`).
		At(m["y"])
}

//doc:summary   Detects useless local constants
//doc:tags      diagnostic
func uselessLocalConst(m dsl.Matcher) {
	m.Match(`const $x = $_; $*body`, `const $x $_ = $_; $*body`).
		Where(!m["x"].Object.IsGlobal() && !m["body"].Contains(`$x`)).
		Report(`useless local constant`)
}

//doc:summary   Detects variables assigment before return that can be simplified
//doc:tags      style
func oneLineReturn(m dsl.Matcher) {
	isPointer := func(x dsl.Var) bool {
		return x.Type.Underlying().Is("*$_") || x.Type.Underlying().Is("chan $_") ||
			x.Type.Underlying().Is("map[$_]$_") || x.Type.Underlying().Is("interface{$*_}") ||
			x.Type.Underlying().Is(`func($*_) $*_`) || x.Type.Underlying().Is(`unsafe.Pointer`)
	}

	m.Match(
		`var $x = $v; return $x`,
		`$x := $v; return $x`,
	).
		Where(!isPointer(m["x"])).
		Suggest(`return $v`)
}
