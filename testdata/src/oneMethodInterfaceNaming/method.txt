package oneMethodInterfaceNaming

type Foo interface { // want `\Qchange interface name to Get + 'er'`
	Get() string
}

type FooBar interface { // want `\Qchange interface name to Pop + 'er'`
	Pop(s string)
}

type FooBar2 interface { // want `\Qchange interface name to Pop + 'er'`
	Pop(s string) string
}

type FooBarFew interface {
	Get() string
	Set()
}

type Replacer interface {
	Replace() string
}
