package unusedLocalConst

func foo() {
	const foo = ""         // want `useless local constant`
	const foo2 string = "" // want `useless local constant`
	{
		const foo3 = "" // want `useless local constant`
	}

	_ = func() {
		const foo4 = 1 // want `useless local constant`
	}
}
