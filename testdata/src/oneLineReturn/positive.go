package oneLineReturn

func myFunc() string                     { return "" }
func myFuncWithArgs(s, ss string) string { return s + ss }

func warning1() string {
	if true {
		var s = "" // want ` suggestion: return \"\"`
		return s
	}
	if true {
		var s = myFunc() // want `suggestion: return myFunc()`
		return s
	}
	if true {
		var s = myFuncWithArgs("12", "34") // want `\Qsuggestion: return myFuncWithArgs("12", "34")`
		return s
	}

	s := "" // want ` suggestion: return \"\"`
	return s
}
