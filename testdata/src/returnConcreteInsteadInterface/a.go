package returnConcreteInsteadInterface

type MyIface interface{}
type s struct{}

func MyFunc() MyIface { // want "in exported functions return concrete type instead of interface"
	return nil
}

func MyFuncWithParams(p, b string) MyIface { // want "in exported functions return concrete type instead of interface"
	return nil
}

func myFunc() MyIface {
	return nil
}

func (s *s) MyFunc() MyIface { // want "in exported functions return concrete type instead of interface"
	return nil
}

func (s *s) myFunc() MyIface {
	return nil
}

func (*s) MyFuncWithoutReceiver() MyIface { // want "in exported functions return concrete type instead of interface"
	return nil
}

func (*s) MyFuncWithoutReceiverWithParams(i bool, k interface{}) MyIface { // want "in exported functions return concrete type instead of interface"
	return nil
}

func (*s) myFuncWithoutReceiver() MyIface {
	return nil
}

func (*s) MyFuncError() error {
	return nil
}

func MyFuncError() error {
	return nil
}

func (s *s) MyFuncErrorWithReceiver() error {
	return nil
}

func MyFuncErrorTwoReturns() (MyIface, error) { //TODO: want "in exported functions return concrete type instead of interface"
	return nil, nil
}

func MyFuncErrorTwoReturnsWithParams(i bool) (MyIface, error) { //TODO: want "in exported functions return concrete type instead of interface"
	return nil, nil
}

func (*s) MyFuncErrorTwoReturns() (MyIface, error) { //TODO: want "in exported functions return concrete type instead of interface"
	return nil, nil
}

func (s *s) MyFuncErrorTwoReturnsWithReceiver() (MyIface, error) { //TODO: want "in exported functions return concrete type instead of interface"
	return nil, nil
}
