package camelCaseNaming

const (
	snake_case_grouped           = "foo" // TODO: want `use camelCase naming strategy`
	snake_case_grouped_2  string = "foo" // TODO: want `use camelCase naming strategy`
	_snake_case_grouped_3        = "foo" // TODO: want `use camelCase naming strategy`
)

const snake_case = "foo"          // want `use camelCase naming strategy`
const _snake_case = "foo"         // want `use camelCase naming strategy`
const _s = "foo"                  // want `use camelCase naming strategy`
const s_ = "foo"                  // want `use camelCase naming strategy`
const snake_case_2 string = "foo" // want `use camelCase naming strategy`

const _ string = "foo"
const _ = "foo"

const (
	_        = "foo"
	_        = "foo"
	_ string = "foo"
)

const (
	_ string = "foo"
)

func foo_bar()                                             {}                 // want `use camelCase naming strategy`
func foo_bar_params(d int)                                 {}                 // want `use camelCase naming strategy`
func foo_bar_fewParams(a, b int)                           {}                 // want `use camelCase naming strategy`
func foo_bar_fewParamsWithReturn(a, b int) (string, error) { return "", nil } // want `use camelCase naming strategy`

type foo struct{}
type foo_bar_struct struct{}           // want `use camelCase naming strategy`
type foo_bar_structWithParams struct { // want `use camelCase naming strategy`
	a string
}

type fooType string
type foo_type string // want `use camelCase naming strategy`

type (
	foo_multiplie_def           struct{}           // TODO: want `use camelCase naming strategy`
	foo_multiplie_defWithParams struct{ a string } // TODO: want `use camelCase naming strategy`
)
