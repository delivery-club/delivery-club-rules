package commonCamelCaseNaming

const (
	snake_case_grouped    = "foo" // want `use camelCase naming strategy`
	snake_case_grouped_2  = "foo" // TODO: want `use camelCase naming strategy`
	_snake_case_grouped_3 = "foo" // TODO: want `use camelCase naming strategy`
)

const (
	snake_case_grouped_with_type    string = "foo" // want `use camelCase naming strategy`
	snake_case_grouped_2_with_type  string = "foo" // TODO: want `use camelCase naming strategy`
	_snake_case_grouped_3_with_type string = "foo" // TODO: want `use camelCase naming strategy`
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

type camelCaseType string
type snake_case_type string // want `use camelCase naming strategy`

type (
	snake_type_grouped   string // TODO: want `use camelCase naming strategy`
	snake_type_grouped_2 string // TODO: want `use camelCase naming strategy`
	camelCaseTypeGrouped string
)

type (
	snake_struct_grouped             struct{}              // TODO: want `use camelCase naming strategy`
	snake_struct_grouped_with_params struct{ a, b string } // TODO: want `use camelCase naming strategy`
	camelCaseStructGrouped           struct{}
	camelCaseStructGroupedWithParams struct{ a, c string }
)

var global_snake_case = 0 // want `use camelCase naming strategy`
var globalCamelCase = 0

func fooFunc(i int) {
	snake_case_var := i             // want `use camelCase naming strategy`
	var snake_case_var_dec = i      // want `use camelCase naming strategy`
	var snake_case_var_full int = i // want `use camelCase naming strategy`
	var (
		snake_case_grouped_var   = i // TODO: want `use camelCase naming strategy`
		snake_case_grouped_var_2 = i // TODO: want `use camelCase naming strategy`
	)

	var (
		snake_case_grouped_var_full   int = i // TODO: want `use camelCase naming strategy`
		snake_case_grouped_var_full_2 int = i // TODO: want `use camelCase naming strategy`
	)

	_ = snake_case_var
	_ = snake_case_var_dec
	_ = snake_case_var_full
	_ = snake_case_grouped_var
	_ = snake_case_grouped_var_2
	_ = snake_case_grouped_var_full
	_ = snake_case_grouped_var_full_2

	camelCaseVar := i
	var camelCaseVarDec = i
	var camelCaseVarFull int = i
	var (
		camelCaseGroupedVar  = i
		camelCaseGroupedVar2 = i
	)

	var (
		camelCaseGroupedVarFull  int = i
		camelCaseGroupedVarFull2 int = i
	)

	_ = camelCaseVar
	_ = camelCaseVarDec
	_ = camelCaseVarFull
	_ = camelCaseGroupedVar
	_ = camelCaseGroupedVar2
	_ = camelCaseGroupedVarFull
	_ = camelCaseGroupedVarFull2
}
