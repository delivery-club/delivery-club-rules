package simplifyErrorCheck

type foo string

func a() {
	var bar foo

	err := bar.Error() // want "error check can be simplified in one line"
	if err != nil {
		return
	}

	err = bar.Error() // want "error check can be simplified in one line"
	if err != nil {
		return
	}

	return
}

func b() {
	var (
		err error
		bar foo
	)

	err = bar.Error() // want "error check can be simplified in one line"
	if err != nil {
		return
	}
}

func myFooFoo() {
	var err = foo("").Error() // want "error check can be simplified in one line"
	if err != nil {
		return
	}

	print(123)
}

func negative() {
	err, err2 := foo("").Error(), foo("").Error()
	if err != nil {
		print(123)
	}

	if err2 != nil {
		print(332)
	}
}

func (foo) Error() error {
	return nil
}
