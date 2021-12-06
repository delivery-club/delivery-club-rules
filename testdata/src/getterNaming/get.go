package getterNaming

type foo string
type bar struct {
	k string
}

func (f foo) GetValue() string     { return string(f) }   // want `don't use 'get' in getter functions`
func (b bar) GetBar() string       { return b.k }         // want `don't use 'get' in getter functions`
func (b bar) GetBarString() string { return string(b.k) } // want `don't use 'get' in getter functions`

func (f *foo) GetValueByPointer() string     { return string(*f) }  // want `don't use 'get' in getter functions`
func (b *bar) GetBarByPointer() string       { return b.k }         // want `don't use 'get' in getter functions`
func (b *bar) GetBarStringByPointer() string { return string(b.k) } // want `don't use 'get' in getter functions`

func (f *foo) nonGetterFunc() string {
	d := ""
	println(d)

	return d
}
