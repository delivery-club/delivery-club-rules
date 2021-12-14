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
	var err error
	if err = myBar(); err != nil { // want `may be simplified to return error without if statement`
		return err
	}

	return nil
}

func myFuncBad() error {
	if true {
		if _, err := myBarSecond(); err != nil { // want `may be simplified to return error without if statement`
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
