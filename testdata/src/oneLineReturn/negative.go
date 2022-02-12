package oneLineReturn

func myFuncError() (string, error) { return "", nil }
func myError() error               { return nil }
func negative() error {
	if true {
		_, err := myFuncError()
		return err
	}
	if true {
		_ = myError()
		return nil
	}

	return nil
}
