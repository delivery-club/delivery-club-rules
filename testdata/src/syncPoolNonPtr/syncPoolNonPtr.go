package syncPoolNonPtr

import (
	"sync"
	"unsafe"
)

type (
	r string
	u struct {
		l []byte
	}
	e struct {
		g string
	}
	uu struct {
		a [102]byte
	}
	rr     []byte
	d      []string
	dd     *d
	aliasD d
)

func foo() {
	var s = sync.Pool{}

	gu := ""
	s.Put(gu) // want `non-pointer values in sync.Pool involve extra allocation`

	bar := r("")

	s.Put(bar) // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put(&bar)

	uv := u{}
	s.Put(uv) // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put(&uv)
	s.Put(u{}) // want `non-pointer values in sync.Pool involve extra allocation`

	ee := e{}
	s.Put(ee)  // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put(e{}) // want `non-pointer values in sync.Pool involve extra allocation`

	uuu := uu{}
	s.Put(uuu) // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put(&uuu)
	s.Put(uu{}) // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put(0)    // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put("")   // want `non-pointer values in sync.Pool involve extra allocation`

	s.Put([]int{123, 213}) // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put(make(chan string))
	s.Put(make(map[int]int, 10))

	s.Put(rr{}) // want `non-pointer values in sync.Pool involve extra allocation`
	s.Put(&rr{})

	var ddObj dd = &d{"123", "1333"}
	s.Put(ddObj)

	s.Put(aliasD{}) // want `non-pointer values in sync.Pool involve extra allocation`
}

func (rec *r) FooBar() {
	k := sync.Pool{}

	k.Put(rec)

	v := unsafe.Pointer(rec)

	k.Put(v)
	k.Put(&v)
}
