package unusedLocalConst

func foo() {
	const foo = ""         // want `unusable local constant`
	const foo2 string = "" // want `unusable local constant`
	{
		const foo3 = "" // want `unusable local constant`
	}

	_ = func() {
		const foo4 = 1 // want `unusable local constant`
	}
}
