package simplifyErrorCheck

type foo string

func a() {
	var bar foo

	err := bar.Error() // want "error check can be simplified in one line"
	if err != nil {
		return
	}

	err = bar.Error()
	if err != nil {
		return
	}

	return
}

func (foo) Error() error {
	return nil
}
