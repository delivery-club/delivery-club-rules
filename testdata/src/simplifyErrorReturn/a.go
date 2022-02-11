package simplifyReturn

func myBar() error {
	return nil
}

func myBarSecond() (string, error) {
	return "", nil
}

func fooBar() error {
	if err := myBar(); err != nil { // want `may be simplified to return error without if statement`
		return err
	}

	return nil
}

func barFoo() error {
	if err := myBar(); err != nil { // want `may be simplified to return error without if statement`
		return err
	}

	return nil
}

func myFuncBad() error {
	if true {
		if _, err := myBarSecond(); err != nil { // want `\Qmay be simplified to return error without if statement`
			return err
		}

		return nil
	}

	{
		if _, err := myBarSecond(); err != nil { // want `may be simplified to return error without if statement`
			return err
		}

		return nil
	}
}

func myFuncGood() (string, error) {
	str, err := myBarSecond()
	if err != nil {
		return "", err
	}

	return str, nil
}

func myFuncSecond() error {
	_, err := myBarSecond() // want `may be simplified to return error without if statement`
	if err != nil {
		return err
	}

	return nil
}

func myFuncThree() error {
	if true {
		var err error

		_, err = myBarSecond() // want `may be simplified to return error without if statement`
		if err != nil {
			return err
		}

		return nil
	}

	var err error
	if err = myBar(); err != nil { // want `may be simplified to return error without if statement`
		return err
	}

	return nil
}
