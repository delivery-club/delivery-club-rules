package rangeExprCopy

func warnings() {
	{
		var xs [777]byte
		for _, x := range xs { // want "copy of xs can be avoided with &xs"
			_ = x
		}
	}

	{
		var foo struct {
			arr [768]byte
		}
		for _, x := range foo.arr { // want "copy of foo.arr can be avoided with &foo.arr"
			_ = x
		}
	}

	{
		xsList := make([][512]byte, 1)
		for _, x := range xsList[0] { // want `\Qcopy of xsList[0] can be avoided with &xsList[0]`
			_ = x
		}
	}

	var x byte
	{
		var xs [777]byte
		for _, x = range xs { // want "copy of xs can be avoided with &xs"
			_ = x
		}
	}

	{
		var foo struct {
			arr [768]byte
		}
		for _, x = range foo.arr { // want "copy of foo.arr can be avoided with &foo.arr"
			_ = x
		}
	}

	{
		xsList := make([][512]byte, 1)
		for _, x = range xsList[0] { // want `\Qcopy of xsList[0] can be avoided with &xsList[0]`
			_ = x
		}
	}
}

func returnArray() [20]int {
	return [20]int{}
}

func noWarnings() {
	// OK: returned value is not addressable, can't take address.
	for _, x := range returnArray() {
		_ = x
	}

	{
		var xs [200]byte
		// OK: already iterating over a pointer.
		for _, x := range &xs {
			_ = x
		}
		// OK: only index is used. No copy is generated.
		for i := range xs {
			_ = xs[i]
		}
		// OK: like in case above, no copy, so it's OK.
		for range xs {
		}
	}

	{
		var xs [10]byte
		// OK: xs is a very small array that can be trivially copied.
		for _, x := range xs {
			_ = x
		}
	}
}
