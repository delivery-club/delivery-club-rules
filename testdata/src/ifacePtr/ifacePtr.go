package ifacePtr

import "fmt"

func ifacePtr() {
	type structType struct {
		_ *fmt.Stringer // want `\Qdon't use pointers to an interface`
	}

	type ifacePtrAlias = *fmt.Stringer // want `\Qdon't use pointers to an interface`

	{
		var x *interface{} // want `\Qdon't use pointers to an interface`
		_ = x
		_ = *x
	}

	{
		var x **interface{} // want `\Qdon't use pointers to an interface`
		_ = x
		_ = *x
	}
}
