package returnConcreteInsteadInterface

type MyIface interface{}
type s struct{}

func MyFunc() MyIface { // want "in exported functions return concrete type instead of interface"
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

func (*s) MyFuncWithoutReciever() MyIface { // want "in exported functions return concrete type instead of interface"
	return nil
}

func (*s) myFuncWithoutReciever() MyIface {
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

func (*s) MyFuncErrorTwoReturns() (MyIface, error) { // TODO: want "in exported functions return concrete type instead of interface"
	return nil, nil
}
