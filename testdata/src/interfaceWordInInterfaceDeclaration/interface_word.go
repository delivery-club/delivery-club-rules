package interfaceWordInInterfaceDeclaration

type interfaceSetter interface { // want `don't use 'interface' word' in interface declaration`
	Get() string
}

type FooInterface interface { // want `don't use 'interface' word' in interface declaration`
	Get() string
	Set() string
}

type FooInterface2 interface { // want `don't use 'interface' word' in interface declaration`
	Get() string
	Set() string
}

type getter interface{ get() string }
