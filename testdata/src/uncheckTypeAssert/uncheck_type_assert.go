package uncheckTypeAssert

func sink(args ...interface{}) {}

func uncheckedTypeAssert() {
	var v interface{}

	_ = v.(int) // want `\Qavoid unchecked type assertions as they can panic`
	{
		x := v.(int) // want `\Qavoid unchecked type assertions as they can panic`
		_ = x
	}

	sink(v.(int))          // want `\Qavoid unchecked type assertions as they can panic`
	sink(0, v.(int))       // want `\Qavoid unchecked type assertions as they can panic`
	sink(v.(int), 0)       // want `\Qavoid unchecked type assertions as they can panic`
	sink(1, 2, v.(int), 3) // want `\Qavoid unchecked type assertions as they can panic`

	{
		type structSink struct {
			f0 interface{}
			f1 interface{}
			f2 interface{}
		}
		_ = structSink{v.(int), 0, 0}    // want `\Qavoid unchecked type assertions as they can panic`
		_ = structSink{0, v.(string), 0} // want `\Qavoid unchecked type assertions as they can panic`
		_ = structSink{0, 0, v.([]int)}  // want `\Qavoid unchecked type assertions as they can panic`

		_ = structSink{f0: v.(int)}                  // want `\Qavoid unchecked type assertions as they can panic`
		_ = structSink{f0: 0, f1: v.(int)}           // want `\Qavoid unchecked type assertions as they can panic`
		_ = structSink{f0: 0, f1: 0, f2: v.(int)}    // want `\Qavoid unchecked type assertions as they can panic`
		_ = structSink{f0: v.(string), f1: 0, f2: 0} // want `\Qavoid unchecked type assertions as they can panic`
	}

	{
		_ = []interface{}{v.(int)}       // want `\Qavoid unchecked type assertions as they can panic`
		_ = []interface{}{0, v.(int)}    // want `\Qavoid unchecked type assertions as they can panic`
		_ = []interface{}{v.(int), 0}    // want `\Qavoid unchecked type assertions as they can panic`
		_ = []interface{}{0, v.(int), 0} // want `\Qavoid unchecked type assertions as they can panic`

		_ = [...]interface{}{10: v.(int)}               // want `\Qavoid unchecked type assertions as they can panic`
		_ = [...]interface{}{10: 0, 20: v.(int)}        // want `\Qavoid unchecked type assertions as they can panic`
		_ = [...]interface{}{10: v.(int), 20: 0}        // want `\Qavoid unchecked type assertions as they can panic`
		_ = [...]interface{}{10: 0, 20: v.(int), 30: 0} // want `\Qavoid unchecked type assertions as they can panic`
	}
}

func negative() {
	var v interface{}

	switch v.(type) {
	case string:
	case int:
	}

	switch vv := v.(type) {
	case string:
		print(vv)
	case byte:
		print(vv)
	}
}
