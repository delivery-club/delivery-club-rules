package rangeCopyVal

type bigObject struct {
	// Fields are carefuly selected to get equal struct size
	// for both AMD64 and 386.

	body [1024]byte
	x    int32
	y    int32
}

func bigCopy(xs []bigObject) int32 {
	v := int32(0)
	for _, x := range xs { // want `\Qeach iteration copies more than 256 bytes (consider pointers or indexing)`
		v += x.x
	}
	return v
}

func bigIndex(xs []bigObject) int32 {
	// OK: no copies.
	v := int32(0)
	for i := range xs {
		v += xs[i].x
	}
	return v
}

func bigTakeAddr(xs []bigObject) int32 {
	// OK: manually taking pointers.
	v := int32(0)
	for i := range xs {
		x := &xs[i]
		v += x.x
	}
	return v
}

func bigPointers(xs []*bigObject) int32 {
	// OK: xs store pointers.
	v := int32(0)
	for _, x := range xs {
		v += x.x
	}
	return v
}
